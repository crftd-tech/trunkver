package ci

func init() {
	RegisterCi(&SimpleEnvBased{
		name:         "Gitlab",
		DetectKey:    "GITLAB_CI",
		SourceRefKey: "CI_COMMIT_SHA",
		BuildRefKey:  "CI_JOB_ID",
		ScmPrefix:    "g",
	})
}
