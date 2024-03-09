package paral

func McloseNil() {
	var c chan string
	close(c)
}
