package basic

import (
	"fmt"
	"os"
	"strconv"
)

func IntString() {
	if len(os.Args) < 3 {
		fmt.Println("Print provide an integer.")
		return
	}

	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Using strconv.Itoa()
	input := strconv.Itoa(n)
	fmt.Printf("strconv.Itoa() %s of type %T\n", input, input)

	// Using strconv.FormatInt
	input = strconv.FormatInt(int64(n), 10)
	fmt.Printf("strconv.FormatInt() %s of type %T\n", input, input)

	// Using string()
	input = string(n)
	fmt.Printf("string() %s of type %T\n", input, input)
}
