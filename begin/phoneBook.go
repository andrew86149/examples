package begin

import (
	"fmt"
	"os"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

func search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func PhoneBook() {
	arguments := os.Args
	if len(arguments) < 3 {
		exe := arguments[0]
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	data = append(data, Entry{"Mihalis", "Tsoukalos", "2109416471"})
	data = append(data, Entry{"Mary", "Doe", "2109416871"})
	data = append(data, Entry{"John", "Black", "2109416123"})

	// Differentiate between the commands
	switch arguments[2] {
	// The search command
	case "search":
		if len(arguments) < 4 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := search(arguments[3])
		if result == nil {
			fmt.Println("Entry not found:", arguments[3])
			return
		}
		fmt.Println(*result)
	// The list command
	case "list":
		list()
	// Anything that is not a match
	default:
		fmt.Println("Not a valid option")
	}
}
