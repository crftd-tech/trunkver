package ci

import (
	"os"
)

type Github struct {
}

// Name implements CI.
func (g *Github) Name() string {
	return "Github"
}

func (g *Github) IsInUse() bool {
	return os.Getenv("GITHUB_SHA") != ""
}

func (g *Github) Get() CIData {
	return CIData{
		SourceRef: "g" + os.Getenv("GITHUB_SHA"),
		BuildRef:  os.Getenv("GITHUB_RUN_ID") + "-" + os.Getenv("GITHUB_RUN_ATTEMPT"),
	}
}

func init() {
	RegisterCi(&Github{})
}
