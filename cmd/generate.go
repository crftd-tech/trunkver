package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"text/template"
	"time"

	"github.com/crftd-tech/trunkver/internal"
	"github.com/crftd-tech/trunkver/internal/ci"
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
		var timestamp string = cmd.Flags().Lookup("timestamp").Value.String()
		var prerelease bool = cmd.Flags().Lookup("prerelease").Value.String() == "true"
		var fileOutput string = cmd.Flags().Lookup("output").Value.String()

		ciResult, found := ci.DetectCi()
		if found {
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

		if sourceRef == "" {
			fmt.Fprintln(os.Stderr, "Error: --source-ref missing, your CI might be unsupported. It should identify the commit that was used to build this artifact, e.g. \"g${GITHUB_SHA:0:7}\" or \"g$(git rev-parse --short HEAD)\".")
			os.Exit(1)
		}

		var parsedTime time.Time
		if timestamp == "now" {
			parsedTime = time.Now()
		} else {
			parsedTime = internal.Must(time.Parse(time.RFC3339, timestamp))
		}

		var trunkVer string

		if prerelease {
			trunkVer = trunkver.GeneratePrereleaseTrunkver(parsedTime, sourceRef, buildRef)
			var baseVersion string
			if len(args) == 0 {
				stdin := cmd.InOrStdin()
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

			if baseVersion != "" {
				trunkVer = trunkver.MergeWithBaseVersion(baseVersion, trunkVer)
			}
		} else {
			trunkVer = trunkver.GenerateMajorTrunkver(parsedTime, sourceRef, buildRef)
		}

		if format := cmd.Flags().Lookup("format").Value.String(); format != "" {
			var tpl = template.Must(template.New("trunkver").Parse(format))
			var buffer bytes.Buffer
			internal.Must(tpl.Execute(&buffer, trunkVer), nil)
			trunkVer = buffer.String()
		}

		fmt.Println(trunkVer)
		if fileOutput != "" {
			internal.Must(os.WriteFile(fileOutput, []byte(trunkVer+"\n"), 0644), nil)
		}
	},
}

func init() {
	generateCmd.Flags().StringP("build-ref", "b", "", "The build ref to use (e.g. $GITHUB_RUN_ID)")
	generateCmd.Flags().StringP("source-ref", "s", "", "The source ref to use for the version (e.g. \"g$(git rev-parse --short HEAD)\")")
	generateCmd.Flags().StringP("timestamp", "t", "now", "The timestamp to use for the version in RFC3339 format")
	generateCmd.Flags().StringP("output", "o", "", "Write TrunkVer to file")
	generateCmd.Flags().StringP("format", "f", "", "Use template to format TrunkVer")
	generateCmd.Flags().BoolP("prerelease", "p", false, "Build the TrunkVer as the prerelease part of a semver (e.g. for nightly builds)")

	rootCmd.AddCommand(generateCmd)
}
