package begin

import (
	"fmt"
)

func Input() {
	// Get User Input
	fmt.Printf("Please give me your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is", name)
}
