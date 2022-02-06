package pkg

func LogIfError(err error) {
	if err != nil {
		panic(err)
	}
}
