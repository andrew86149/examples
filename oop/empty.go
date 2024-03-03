package oop

import "fmt"

type S11 struct {
	F1 int
	F2 string
}

type S21 struct {
	F1 int
	F2 S11
}

func Print(s interface{}) {
	fmt.Println(s)
}

func Mempty() {
	v1 := S11{10, "Hello"}
	v2 := S21{F1: -1, F2: v1}
	Print(v1)
	Print(v2)
	// Printing an integer
	Print(123)
	// Printing a string
	Print("Go is the best!")
}
