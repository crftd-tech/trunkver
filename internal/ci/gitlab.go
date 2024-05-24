package ci

import (
	"os"
)

type Gitlab struct{}

func (g *Gitlab) IsInUse() bool {
	return os.Getenv("GITLAB_CI") == "true"
}

func (g *Gitlab) Get() CIData {
	return CIData{
		SourceRef: "g" + os.Getenv("CI_COMMIT_SHA")[:7],
		BuildRef:  os.Getenv("CI_JOB_ID"),
	}
}

func init() {
	RegisterCi(&Gitlab{})
}
