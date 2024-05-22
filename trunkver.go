package main

import (
	"time"
)

func formatTrunkver(ts time.Time, sourceRef, buildRef string) string {
	return ts.UTC().Format("20060102150405") + ".0.0-" + sourceRef + "-" + buildRef
}
