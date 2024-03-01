package complexn

import (
	"fmt"
)

func NilMap() {
	aMap := map[string]int{}
	aMap["test"] = 1
	fmt.Println("aMap:", aMap)
	aMap = nil
	fmt.Println("aMap:", aMap)
	if aMap == nil {
		fmt.Println("nil map!")
		aMap = map[string]int{}
	}
	aMap["test"] = 10
	fmt.Println("aMap:", aMap)
	// This will crash!
	aMap = nil
	fmt.Println("aMap:", aMap)
	if aMap == nil {
		fmt.Println("nil map!")
		aMap = map[string]int{}
	}
	aMap["test"] = 100
	fmt.Println("aMap:", aMap)
	aMap = nil

	//aMap["test"] = 1000
	fmt.Println("aMap:", aMap)
}
