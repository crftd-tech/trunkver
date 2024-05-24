package ci

import (
	"os"
)

type Github struct{}

func (g *Github) IsInUse() bool {
	return os.Getenv("GITHUB_SHA") != "" && os.Getenv("GITHUB_RUN_ID") != ""
}

func (g *Github) Get() CIData {
	return CIData{
		SourceRef: "g" + os.Getenv("GITHUB_SHA")[:7],
		BuildRef:  os.Getenv("GITHUB_RUN_ID"),
	}
}

func init() {
	RegisterCi(&Github{})
}
