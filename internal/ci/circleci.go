package ci

import (
	"os"
)

type Circleci struct{}

func (g *Circleci) IsInUse() bool {
	return os.Getenv("CIRCLECI") == "true"
}

func (g *Circleci) Get() CIData {
	return CIData{
		SourceRef: "g" + os.Getenv("CIRCLE_SHA1")[:7],
		BuildRef:  os.Getenv("CIRCLE_WORKFLOW_JOB_ID"),
	}
}

func init() {
	RegisterCi(&Circleci{})
}
