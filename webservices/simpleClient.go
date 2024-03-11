package webservices

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func MsimpleClient() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}

	URL := os.Args[2]
	data, err := http.Get(URL)

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = io.Copy(os.Stdout, data.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	data.Body.Close()
}