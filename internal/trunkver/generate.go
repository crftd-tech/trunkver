package trunkver

import (
	"regexp"
	"strconv"
	"time"

	"github.com/Masterminds/semver/v3"
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
	major, err := strconv.ParseUint(ts.UTC().Format("20060102150405"), 10, 64)
	if err != nil {
		panic(err)
	}
	var sVer semver.Version = *semver.New(
		major, 0, 0,
		coerceSourceRef(sourceRef)+"-"+coerceBuildRef(buildRef),
		"",
	)

	trunkVer := sVer.String()
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
