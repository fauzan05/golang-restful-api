package helper

func HandleErrorWithPanic(err error) {
	if err != nil {
		panic(err.Error())
	}
}