package trunkver

import (
	"github.com/Masterminds/semver/v3"
)

func MergeWithBaseVersion(baseVersion string, prereleaseTrunkVer string) string {
	var semverBaseVersion, err = semver.NewVersion(baseVersion)
	if err != nil {
		panic(err)
	}

	var newVersion semver.Version
	newVersion, err = semverBaseVersion.SetPrerelease(prereleaseTrunkVer)
	if err != nil {
		panic(err)
	}
	return newVersion.String()
}
