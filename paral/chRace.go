package paral

import (
	"fmt"
)

func printer2(ch chan<- bool, times int) {
	for i := 0; i < times; i++ {
		ch <- true
	}
	close(ch)
}

func MchRace() {
	// This is an unbuffered channel
	var ch chan bool = make(chan bool)

	// Write 5 values to channel with a single goroutine
	go printer2(ch, 5)

	// IMPORTANT: As the channel c is closed,
	// the range loop is going to exit on its own.
	for val := range ch {
		fmt.Print(val, " ")
	}
	fmt.Println()

	for i := 0; i < 15; i++ {
		fmt.Print(<-ch, " ")
	}
	fmt.Println()
}
