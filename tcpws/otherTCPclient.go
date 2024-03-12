package tcpws

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func MtcpCO() {
	arguments := os.Args
	if len(arguments) == 2 {
		fmt.Println("Please provide a server:port string!")
		return
	}

	connect := arguments[2]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", connect)
	if err != nil {
		fmt.Println("ResolveTCPAddr:", err)
		return
	}

	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		fmt.Println("DialTCP:", err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			conn.Close()
			return
		}
	}
}
