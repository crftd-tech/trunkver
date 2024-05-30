package internal

func Must[E any](e E, err error) E {
	if err != nil {
		panic(err)
	}
	return e
}

func Must2[E any, F any](e E, f F, err error) (E, F) {
	if err != nil {
		panic(err)
	}
	return e, f
}
