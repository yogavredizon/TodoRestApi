package exception

func Validation(err error) {
	if err != nil {
		panic(err)
	}
}
