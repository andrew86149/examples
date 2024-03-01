package complexn

import (
	"fmt"
	"os"
	"regexp"
)

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func NameSurRE() {
	arguments := os.Args
	if len(arguments) < 3 {
		fmt.Println("Usage: <utility> string.")
		return
	}

	s := arguments[2]
	ret := matchNameSur(s)
	fmt.Println(ret)
}
