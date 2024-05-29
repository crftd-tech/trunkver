package ci

func init() {
	RegisterCi(&SimpleEnvBased{
		Name:         "Gitlab",
		DetectKey:    "GITLAB_CI",
		SourceRefKey: "CI_COMMIT_SHA",
		BuildRefKey:  "CI_JOB_ID",
		ScmPrefix:    "g",
	})
}
