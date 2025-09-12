package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"text/template"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/crftd-tech/trunkver/internal"
	"github.com/crftd-tech/trunkver/internal/ci"
	"github.com/crftd-tech/trunkver/internal/log"
	"github.com/crftd-tech/trunkver/internal/trunkver"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate [flags] [base-version | -]",
	Aliases: []string{"gen", "g"},
	Short:   "Generate a new TrunkVer",
	Long:    `Generates a new TrunkVer, optionally appending it to base-version as the prerelease part (if --prerelease).`,
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var buildRef string = cmd.Flags().Lookup("build-ref").Value.String()
		var sourceRef string = cmd.Flags().Lookup("source-ref").Value.String()
		var sourceRefMaxLength, _ = cmd.Flags().GetInt("source-ref-max-length")
		var timestamp string = cmd.Flags().Lookup("timestamp").Value.String()
		var prerelease bool = cmd.Flags().Lookup("prerelease").Value.String() == "true"
		var incrementBaseVersionPart string = cmd.Flags().Lookup("increment").Value.String()
		var fileOutput string = cmd.Flags().Lookup("output").Value.String()

		ciResult, found := ci.DetectCi()
		if found {
			log.LogVerbose("CI detected: %s", ciResult.Name())
			ciData := ciResult.Get()
			if sourceRef == "" {
				sourceRef = ciData.SourceRef
			}
			if buildRef == "" {
				buildRef = ciData.BuildRef
			}
		}

		if buildRef == "" {
			fmt.Fprintln(os.Stderr, "Error: --build-ref missing, your CI might be unsupported. It should identify the log that was produced during creation of this artifact, e.g. the job id in Github Actions")
			os.Exit(1)
		}

		if sourceRefMaxLength > 0 {
			sourceRef = sourceRef[0:min(len(sourceRef), sourceRefMaxLength)]
		}

		if sourceRef == "" {
			fmt.Fprintln(os.Stderr, "Error: --source-ref missing, your CI might be unsupported. It should identify the commit that was used to build this artifact, e.g. \"g${GITHUB_SHA:0:7}\" or \"g$(git rev-parse --short HEAD)\".")
			os.Exit(1)
		}

		var parsedTime = parseRFC3339OrNow(timestamp)

		var trunkVer string

		if prerelease {
			trunkVer = trunkver.GeneratePrereleaseTrunkver(parsedTime, sourceRef, buildRef)
			var baseVersion = readBaseVersionFromArgsOrStdin(cmd.InOrStdin(), args)
			if baseVersion != "" {
				baseVersion = removeVPrefixIfPresent(baseVersion)
				if incrementBaseVersionPart != "" {
					baseVersion = incBaseVersion(baseVersion, incrementBaseVersionPart)
				}
				trunkVer = trunkver.MergeWithBaseVersion(baseVersion, trunkVer)
			}
		} else {
			trunkVer = trunkver.GenerateMajorTrunkver(parsedTime, sourceRef, buildRef)
		}

		if format := cmd.Flags().Lookup("format").Value.String(); format != "" {
			trunkVer = templateTrunkVer(trunkVer, format)
		}

		if fileOutput != "" {
			log.LogVerbose("Writing %s to %s", trunkVer, fileOutput)
			internal.Must(os.WriteFile(fileOutput, []byte(trunkVer+"\n"), 0644), nil)
		} else {
			fmt.Println(trunkVer)
		}
	},
}

func removeVPrefixIfPresent(baseVersion string) string {
	if baseVersion[0] == 'v' {
		return baseVersion[1:]
	}
	return baseVersion
}
func incBaseVersion(baseVersion string, inc string) string {
	var semverBaseVersion, err = semver.NewVersion(baseVersion)
	if err != nil {
		panic(err)
	}

	switch inc {
	case "major":
		return semverBaseVersion.IncMajor().String()
	case "minor":
		return semverBaseVersion.IncMinor().String()
	case "patch":
		return semverBaseVersion.IncPatch().String()
	default:
		panic("Can't increment " + inc)
	}
}

func init() {
	generateCmd.Flags().StringP("build-ref", "b", "", "The build ref to use (e.g. $GITHUB_RUN_ID)")
	generateCmd.Flags().StringP("source-ref", "s", "", "The source ref to use for the version (e.g. \"g$(git rev-parse --short HEAD)\")")
	generateCmd.Flags().IntP("source-ref-max-length", "", 8, "The length to truncate the source-ref to. Set to 0 to disable truncating.")
	generateCmd.Flags().StringP("timestamp", "t", "now", "The timestamp to use for the version in RFC3339 format")
	generateCmd.Flags().StringP("output", "o", "", "Write TrunkVer to file")
	generateCmd.Flags().StringP("format", "f", "", "Use template to format TrunkVer")
	generateCmd.Flags().StringP("increment", "i", "", "increment the specified version part when generating a prerelease with a given base version (can be patch, minor, major)")
	generateCmd.Flags().BoolP("prerelease", "p", false, "Build the TrunkVer as the prerelease part of a semver (e.g. for nightly builds)")

	rootCmd.AddCommand(generateCmd)
}

func parseRFC3339OrNow(timestamp string) time.Time {
	if timestamp == "now" {
		return time.Now()
	}
	return internal.Must(time.Parse(time.RFC3339, timestamp))

}

func readBaseVersionFromArgsOrStdin(stdin io.Reader, args []string) string {
	var baseVersion string
	if len(args) == 1 {
		if args[0] == "-" {
			reader := bufio.NewReader(stdin)
			inputBA, _, err := reader.ReadLine()
			if err == nil {
				baseVersion = string(inputBA)
			} else if err != io.EOF {
				panic(err)
			}
		} else {
			baseVersion = args[0]
		}
	}
	return baseVersion
}

func templateTrunkVer(trunkVer, format string) string {
	var tpl = template.Must(template.New("trunkver").Parse(format))
	var buffer bytes.Buffer
	internal.Must(tpl.Execute(&buffer, trunkVer), nil)
	return buffer.String()
}
