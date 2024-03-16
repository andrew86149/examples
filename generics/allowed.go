package generics

import (
	"fmt"
)

func Same[T comparable](a, b T) bool {
	//if a == b {
	//	return true
	//}
	//return false
	return a == b
}

func Mallowed() {
	fmt.Println("4 = 3 is", Same(4, 3))
	fmt.Println("aa = aa is", Same("aa", "aa"))
	fmt.Println("4.1 = 4.15 is", Same(4.1, 4.15))
	fmt.Println("4 = 4 is", Same(4, 4))
	fmt.Println("aa = ab is", Same("aa", "ab"))
	fmt.Println("4.1 = 4.1 is", Same(4.1, 4.1))
}
