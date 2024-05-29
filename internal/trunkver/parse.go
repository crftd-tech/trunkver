package trunkver

import (
	"strconv"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
)

type SourceRef struct {
	CommitHash string `json:"commitHash"`
	ScmPrefix  string `json:"scmPrefix"`
	RawRef     string `json:"rawRef"`
}

type TrunkVer struct {
	Timestamp time.Time `json:"timestamp"`
	SourceRef SourceRef `json:"sourceRef"`
	BuildRef  string    `json:"buildRef"`
}

func parseSourceRef(input string) SourceRef {
	return SourceRef{
		CommitHash: input[1:],
		ScmPrefix:  string(input[0]),
		RawRef:     input,
	}
}

func ParseTrunkVer(input string) (*TrunkVer, error) {
	ver, err1 := semver.NewVersion(input)
	if err1 != nil {
		return nil, err1
	}

	ts, err2 := time.Parse("20060102150405", strconv.FormatUint(ver.Major(), 10))
	if err2 != nil {
		return tryParsePreelaseVersion(ver)
	}

	return tryParseMajorVersion(ver, ts)
}

func tryParsePreelaseVersion(ver *semver.Version) (*TrunkVer, error) {
	prereleaseParts := strings.Split(ver.Prerelease(), "-")

	ts, err := time.Parse("20060102150405", prereleaseParts[0])
	if err != nil {
		return nil, err
	}

	return &TrunkVer{
		Timestamp: ts,
		SourceRef: parseSourceRef(prereleaseParts[1]),
		BuildRef:  strings.Join(prereleaseParts[2:], "-"),
	}, nil
}

func tryParseMajorVersion(ver *semver.Version, time time.Time) (*TrunkVer, error) {
	prereleaseParts := strings.Split(ver.Prerelease(), "-")
	return &TrunkVer{
		Timestamp: time,
		SourceRef: parseSourceRef(prereleaseParts[0]),
		BuildRef:  strings.Join(prereleaseParts[1:], "-"),
	}, nil
}
