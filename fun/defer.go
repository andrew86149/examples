package fun

import (
	"fmt"
)

func d1() {
	for i := 3; i > 0; i-- {
		fmt.Print(i, " ")
		defer fmt.Print(i, " ")
	}
}

func d2() {
	fmt.Println()
	for i := 3; i > 0; i-- {
		fmt.Print(i, " ")
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
}

func d3() {
	for i := 3; i > 0; i-- {
		fmt.Print(i, " ")
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}

func Mdefer() {
	d1()
	d2()
	fmt.Println()
	d3()
	fmt.Println()
}
