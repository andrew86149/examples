package paral

import (
	"fmt"
	"time"
)

func MgoClosure() {
	for i := 0; i <= 20; i++ {
		go func() {
			fmt.Print(i, " ")
		}()
	}
	time.Sleep(time.Second)
	fmt.Println()
}
