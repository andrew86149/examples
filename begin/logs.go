package begin

import (
	"log"
	"os"
)

func Logs() {
	if len(os.Args) != 2 {
		log.Fatal("Fatal: Hello World!")
	}
	log.Panic("Panic: Hello World!")
}
