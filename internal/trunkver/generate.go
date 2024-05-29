package trunkver

import "time"

func GenerateMajorTrunkver(ts time.Time, sourceRef, buildRef string) string {
	return ts.UTC().Format("20060102150405") + ".0.0-" + sourceRef + "-" + buildRef
}

func GeneratePrereleaseTrunkver(ts time.Time, sourceRef, buildRef string) string {
	return ts.UTC().Format("20060102150405") + "-" + sourceRef + "-" + buildRef
}
