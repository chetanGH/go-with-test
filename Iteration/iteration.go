package iteration

func Iteration(str string) string {
	var blank string
	for i := 0; i < 5; i++ {
		blank = blank + str
	}
	return blank
}
