package main

import (
	"fmt"
	"net"
)

var manifest map[string]entry

func main() {
	manifest = make(map[string]entry)
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
		go broadcastMessage(string(buffer[:n]), username)
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
	AddUser(string(buffer[:n]), conn, manifest)
	return string(buffer[:n])
}

func broadcastMessage(msg string, username string) {

	for _, v := range manifest {
		data := username + "___" + msg
		_, err := (v.Conn).Write([]byte(data))
		checkerror(err)
	}
}
