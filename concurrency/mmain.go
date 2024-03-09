package concurrency

import "fmt"

func Mmain() {
	fmt.Println("main function started")
	CalcStoreTotal(Products)
	fmt.Println("main function complete")
	//info()
	//init2()
}

/*
func info() {
	for _, p := range ProductList {
		_, ok := Products[p.Category]
		if ok {
			fmt.Println("Products ", p)
		} else {
			fmt.Println("ProductGroup ", p)
		}
	}
}*/
