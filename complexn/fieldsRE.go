package complexn

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

/*
func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}*/

func matchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`\d+$`)
	return re.Match(t)
}

func matchRecord(s string) bool {
	fields := strings.Split(s, ",")
	if len(fields) != 3 {
		return false
	}

	if !matchNameSur(fields[0]) {
		return false
	}

	if !matchNameSur(fields[1]) {
		return false
	}

	return matchTel(fields[2])
}

func FieldsRE() {
	arguments := os.Args
	if len(arguments) < 3 {
		fmt.Println("Usage: <utility> record.")
		return
	}

	s := arguments[2]
	err := matchRecord(s)
	fmt.Println(err)
}
