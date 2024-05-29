package trunkver

import (
	"regexp"
	"time"
)

var sourceRefFilterRegex *regexp.Regexp = regexp.MustCompile(`[^a-zA-Z0-9]`)

func coerceSourceRef(sourceRef string) string {
	return sourceRefFilterRegex.ReplaceAllString(sourceRef, "")
}

var buildRefFilterRegex *regexp.Regexp = regexp.MustCompile(`[^a-zA-Z0-9-]`)

func coerceBuildRef(buildRef string) string {
	return buildRefFilterRegex.ReplaceAllString(buildRef, "-")
}

func GenerateMajorTrunkver(ts time.Time, sourceRef, buildRef string) string {
	trunkVer := ts.UTC().Format("20060102150405") + ".0.0-" + coerceSourceRef(sourceRef) + "-" + coerceBuildRef(buildRef)
	if _, err := ParseTrunkVer(trunkVer); err != nil {
		panic(err)
	}
	return trunkVer
}

func GeneratePrereleaseTrunkver(ts time.Time, sourceRef, buildRef string) string {
	trunkVer := ts.UTC().Format("20060102150405") + "-" + coerceSourceRef(sourceRef) + "-" + coerceBuildRef(buildRef)
	if _, err := ParseTrunkVer(trunkVer); err != nil {
		panic(err)
	}
	return trunkVer
}
