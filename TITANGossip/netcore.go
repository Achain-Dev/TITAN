package main

import (
	"net"
	"golang.org/x/crypto/ssh"
)

// 48563

func WaitForConnections() {


}

func TimeoutConnection(Done chan bool, nConn net.Conn) {

}

func HandleIncomingConn(nConn net.Conn, config *ssh.ServerConfig, IsUserAllowedKeyAuth map[string]bool) {


}

func LoadPrivKeyFromFile(file string) []byte {
	return nil;
}


func ConnectToPeer(P *Peer) error {
	return nil;
}
