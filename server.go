package main

import (
	"fmt"
	"net"
)

func main() {
	serverIP := ":7000"
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverIP)
	checkerror(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	defer listener.Close()
	checkerror(err)
	fmt.Println("Ready to Accept")
	for {
		conn, err := listener.AcceptTCP()
		checkerror(err)
		fmt.Println("New connection")
		go handleRequest(conn)
	}
}

func checkerror(err error) {
	if err != nil {
		fmt.Println("error reported in: ", err)
	}
}

func handleRequest(conn *net.TCPConn) {
	username := extractUsername(conn)
	if username == "" {
		return
	}
	buffer := make([]byte, 128)
	for {
		n, err := conn.Read(buffer)
		if n == 0 {
			break
		}
		checkerror(err)
		fmt.Println(username, " : ", string(buffer[:n]))
	}
}

func extractUsername(conn *net.TCPConn) string {
	buffer := make([]byte, 128)
	var err error
	n, err := conn.Read([]byte(buffer))
	if n == 0 {
		return ""
	}
	checkerror(err)
	fmt.Println("New user entered chatroom : ", string(buffer[:n]))
	return string(buffer[:n])
}
