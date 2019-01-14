package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	serverIP := "localhost:7000"
	reader := bufio.NewReader(os.Stdin)
	var message string
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverIP)
	checkerror(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkerror(err)
	fmt.Println("Please enter Username")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSuffix(username, "\n")
	_, err = conn.Write([]byte(username))
	checkerror(err)
	fmt.Println("You are now connected to the server")
	for {
		message, _ = reader.ReadString('\n')
		message = strings.TrimSuffix(message, "\n")
		_, err = conn.Write([]byte(message))
		checkerror(err)
	}

	conn.Close()
}

func checkerror(err error) {
	if err != nil {
		fmt.Println("error reported in: ", err)
	}
}
