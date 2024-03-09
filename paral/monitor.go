package paral

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var readValue = make(chan int)
var writeValue = make(chan int)

func set(newValue int) {
	writeValue <- newValue
}

func read2() int {
	return <-readValue
}

func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Printf("%d ", value)
		case readValue <- value:
		}
	}
}

func Mmonitor() {
	if len(os.Args) != 3 {
		fmt.Println("Please give an integer!")
		return
	}
	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Going to create %d random numbers.\n", n)
	rand.Seed(time.Now().Unix())
	go monitor()
	var wg sync.WaitGroup

	for r := 0; r < n; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			set(rand.Intn(10 * n))
		}()
	}
	wg.Wait()
	fmt.Printf("\nLast value: %d\n", read2())
}
