package generics

import (
	"fmt"
)

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func Mhw() {
	PrintSlice([]int{1, 2, 3, 4, 5, 6})
	PrintSlice([]string{"a", "b", "c"})
	PrintSlice([]float64{1.2, -2.33, 4.55})
}
