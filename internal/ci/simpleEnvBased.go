package ci

import "os"

type SimpleEnvBased struct {
	Name         string
	ScmPrefix    string
	DetectKey    string
	SourceRefKey string
	BuildRefKey  string
}

func (g *SimpleEnvBased) IsInUse() bool {
	return os.Getenv(g.DetectKey) != ""
}

func (g *SimpleEnvBased) Get() CIData {
	return CIData{
		SourceRef: g.ScmPrefix + os.Getenv(g.SourceRefKey)[:7],
		BuildRef:  os.Getenv(g.BuildRefKey),
	}
}
