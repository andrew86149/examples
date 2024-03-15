package tsts

import (
	"fmt"
	"runtime"
)

func McrossCompile() {
	fmt.Print("You are using ", runtime.GOOS, " ")
	fmt.Println("on a(n)", runtime.GOARCH, "machine")
	fmt.Println("with Go version", runtime.Version())
}
