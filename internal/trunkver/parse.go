package trunkver

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
)

type SourceRef struct {
	CommitHash string `json:"commitHash,omitempty"`
	ScmPrefix  string `json:"scmPrefix,omitempty"`
	RawRef     string `json:"rawRef"`
}

type TrunkVer struct {
	Timestamp time.Time `json:"timestamp"`
	SourceRef SourceRef `json:"sourceRef"`
	BuildRef  string    `json:"buildRef"`
}

func parseSourceRef(input string) SourceRef {
	if input[0] == 'g' {
		return SourceRef{
			CommitHash: input[1:],
			ScmPrefix:  string(input[0]),
			RawRef:     input,
		}
	}
	return SourceRef{
		RawRef: input,
	}
}

func ParseTrunkVer(input string) (*TrunkVer, error) {
	ver, err1 := semver.NewVersion(input)
	if err1 != nil {
		return nil, err1
	}

	ts, err2 := time.Parse("20060102150405", strconv.FormatUint(ver.Major(), 10))
	if err2 != nil {
		return tryParsePrerelaseVersion(ver)
	}

	return tryParseMajorVersion(ver, ts)
}

func tryParsePrerelaseVersion(ver *semver.Version) (*TrunkVer, error) {
	var parts = strings.SplitN(ver.Prerelease(), "-", 3)
	if len(parts) != 3 {
		return nil, fmt.Errorf("PRERELEASE TrunkVer does not contain timestamp, source and build ref in %s", ver.String())
	}

	var timestampPart, sourceRefPart, buildRefPart = parts[0], parts[1], parts[2]

	ts, err := time.Parse("20060102150405", timestampPart)
	if err != nil {
		return nil, err
	}

	return &TrunkVer{
		Timestamp: ts,
		SourceRef: parseSourceRef(sourceRefPart),
		BuildRef:  buildRefPart,
	}, nil
}

func tryParseMajorVersion(ver *semver.Version, time time.Time) (*TrunkVer, error) {
	var parts = strings.SplitN(ver.Prerelease(), "-", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("MAJOR TrunkVer does not contain source and build ref in %s", ver.String())
	}
	var sourceRefPart, buildRefPart = parts[0], parts[1]

	return &TrunkVer{
		Timestamp: time,
		SourceRef: parseSourceRef(sourceRefPart),
		BuildRef:  buildRefPart,
	}, nil
}
