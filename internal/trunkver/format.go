package trunkver

import "time"

func FormatMajorTrunkver(ts time.Time, sourceRef, buildRef string) string {
	return ts.UTC().Format("20060102150405") + ".0.0-" + sourceRef + "-" + buildRef
}

func FormatPrereleaseTrunkver(ts time.Time, sourceRef, buildRef string) string {
	return ts.UTC().Format("20060102150405") + "-" + sourceRef + "-" + buildRef
}
