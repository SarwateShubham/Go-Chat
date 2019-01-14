package main

import "net"

type entry struct {
	Conn *net.TCPConn
}

func AddUser(username string, Conn *net.TCPConn, manifest map[string]entry) {
	manifest[username] = entry{Conn}
}

func DeleteUser(username string, Conn *net.TCPConn, manifest map[string]entry) {
	delete(manifest, username)
}
