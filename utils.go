package main

import "net"

type entry struct {
	Conn *net.TCPConn
}

/*
AddUser provides the functionality to add a user to the manifest consisting of the clients connected to the server.
*/
func AddUser(username string, Conn *net.TCPConn, manifest map[string]entry) {
	manifest[username] = entry{Conn}
}

func DeleteUser(username string, Conn *net.TCPConn, manifest map[string]entry) {
	delete(manifest, username)
}
