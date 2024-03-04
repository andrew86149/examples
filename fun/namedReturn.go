package fun

import (
	"fmt"
	"os"
	"strconv"
)

func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
		return min, max
	}

	min = x
	max = y
	return
}

func MnamedReturn() {
	arguments := os.Args
	if len(arguments) < 4 {
		fmt.Println("The program needs at least 2 arguments!")
		return
	}

	// No checking here - we trust the user!!
	a1, _ := strconv.Atoi(arguments[2])
	a2, _ := strconv.Atoi(arguments[3])

	fmt.Println(minMax(a1, a2))
	mi, ma := minMax(a1, a2)
	fmt.Println(mi, ma)
}
