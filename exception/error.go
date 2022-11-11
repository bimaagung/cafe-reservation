package exception

func Error(err interface{}) {
	if err != nil {
		panic(err)
	}
}