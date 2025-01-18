package ci

type CI interface {
	Name() string
	IsInUse() bool
	Get() CIData
}

type CIData struct {
	SourceRef string
	BuildRef  string
}

// List of ci providers, not a map
var providers []CI

func RegisterCi(ci CI) {
	providers = append(providers, ci)
}

func DetectCi() (CI, bool) {
	for _, ci := range providers {
		if ci.IsInUse() {
			return ci, true
		}
	}
	return nil, false
}
