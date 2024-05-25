package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/crftd-tech/trunkver/internal/ci"
	"github.com/crftd-tech/trunkver/internal/trunkver"
)

var Version string = "0.0.0-HEAD-local"

func main() {
	run(os.Stdout, os.Stderr, os.Args)
}

func run(out io.Writer, err io.Writer, args []string) {
	flagSet := flag.NewFlagSet("trunkver", flag.ExitOnError)
	version := flagSet.Bool("version", false, "Print the version and exit")
	ts := flagSet.String("timestamp", "now", "The timestamp to use for the version in RFC3339 format")
	sRef := flagSet.String("source-ref", "", "The source ref to use for the version")
	bRef := flagSet.String("build-ref", "", "The build ref to use for the version")
	prereleaseOnly := flagSet.Bool("prerelease", false, "Build the trunkver as the prerelease part of a version (e.g. for nightly builds)")

	flagSet.Parse(args[1:])
	baseVersion := flagSet.Arg(0)

	if *version {
		fmt.Fprintln(err, Version)
		return
	}

	ciResult, found := ci.DetectCi()
	if found {
		ciData := ciResult.Get()
		if *sRef == "" {
			*sRef = ciData.SourceRef
		}
		if *bRef == "" {
			*bRef = ciData.BuildRef
		}
	}

	if *bRef == "" {
		fmt.Fprintln(err, "Error: --build-ref missing, your CI might be unsupported. Specify it manually. It should identify the log that was produced during creation of this artifact, e.g. the Job Id in Github Actions")
		os.Exit(1)
	}

	if *sRef == "" {
		fmt.Fprintln(err, "Error: --source-ref missing, your CI might be unsupported. Specify it manually. It should identify the commit that was used to build this artifact, e.g. \"g${GITHUB_SHA:0:7}\" or \"g$(git rev-parse --short HEAD)\".")
		os.Exit(1)
	}

	var parsedTime time.Time
	if *ts == "now" {
		parsedTime = time.Now()
	} else {
		var err error
		parsedTime, err = time.Parse(time.RFC3339, *ts)
		if err != nil {
			panic(err)
		}
	}

	if !*prereleaseOnly {
		fmt.Fprintln(out, trunkver.FormatMajorTrunkver(parsedTime, *sRef, *bRef))
		return
	}

	var trunkVer string = trunkver.FormatPrereleaseTrunkver(parsedTime, *sRef, *bRef)
	if baseVersion == "" {
		fmt.Fprintln(out, trunkVer)
		return
	}

	fmt.Fprintln(out, trunkver.MergeWithBaseVersion(baseVersion, trunkVer))
}
