package tcpws

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func MtcpC() {
	arguments := os.Args
	if len(arguments) == 2 {
		fmt.Println("Please provide host:port.")
		return
	}

	connect := arguments[2]
	c, err := net.Dial("tcp", connect)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
