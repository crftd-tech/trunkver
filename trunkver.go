package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func formatTrunkver(ts time.Time, sourceRef, buildRef string) string {
	return ts.UTC().Format("20060102150405") + ".0.0-g" + sourceRef + "-" + buildRef
}

func main() {
	run(os.Stdout, os.Args)
}

func run(w io.Writer, args []string) {
	flagSet := flag.NewFlagSet("trunkver", flag.ExitOnError)
	ts := flagSet.String("timestamp", "now", "The timestamp to use for the version in RFC3339 format")
	sRef := flagSet.String("source-ref", "", "The source ref to use for the version")
	bRed := flagSet.String("build-ref", "", "The build ref to use for the version")

	flagSet.Parse(args[1:])

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

	fmt.Fprintln(w, formatTrunkver(parsedTime, *sRef, *bRed))
}
