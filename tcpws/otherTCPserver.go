package tcpws

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func MtcpSO() {
	arguments := os.Args
	if len(arguments) == 2 {
		fmt.Println("Please provide a port number!")
		return
	}

	SERVER := "localhost" + ":" + arguments[2]
	s, err := net.ResolveTCPAddr("tcp", SERVER)
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := net.ListenTCP("tcp", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := make([]byte, 1024)
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting TCP server!")
			conn.Close()
			return
		}

		fmt.Print("> ", string(buffer[0:n-1]), "\n")
		_, err = conn.Write(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
