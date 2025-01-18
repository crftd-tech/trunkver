package ci

func init() {
	RegisterCi(&SimpleEnvBased{
		name:         "CircleCI",
		DetectKey:    "CIRCLECI",
		SourceRefKey: "CIRCLE_SHA1",
		BuildRefKey:  "CIRCLE_WORKFLOW_JOB_ID",
		ScmPrefix:    "g",
	})
}
