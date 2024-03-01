package complexn

import (
	"fmt"
	"os"
	"regexp"
)

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func IntRE() {
	arguments := os.Args
	if len(arguments) < 3 {
		fmt.Println("Usage: <utility> string.")
		return
	}

	s := arguments[2]
	ret := matchInt(s)
	fmt.Println(ret)
}
