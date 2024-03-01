package begin

import (
	"fmt"
	"time"
)

func myPrint(start, finish int) {
	for i := start; i <= finish; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	time.Sleep(100 * time.Microsecond)
}

func GoRoutines() {
	for i := 0; i < 4; i++ {
		go myPrint(i, 5)
	}
	time.Sleep(time.Second)
}
