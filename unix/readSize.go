package unix

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func readSize(f *os.File, size int) []byte {
	buffer := make([]byte, size)
	n, err := f.Read(buffer)

	// io.EOF is a special case and is treated as such
	if err == io.EOF {
		return nil
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return buffer[0:n]
}

func MreadSize() {
	arguments := os.Args
	if len(arguments) < 4 {
		fmt.Println("<buffer size> <filename>")
		return
	}

	bufferSize, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	file := os.Args[3]
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	readData := readSize(f, bufferSize)
	if readData != nil {
		fmt.Print(string(readData))
	} else {
		return
	}
	fmt.Println()
}
