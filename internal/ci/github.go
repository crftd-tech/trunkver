package ci

func init() {
	RegisterCi(&SimpleEnvBased{
		Name:         "Github",
		DetectKey:    "GITHUB_SHA",
		SourceRefKey: "GITHUB_SHA",
		BuildRefKey:  "GITHUB_RUN_ID",
	})
}
