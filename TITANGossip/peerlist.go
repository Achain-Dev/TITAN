package main

import (
	"code.google.com/p/go.crypto/ssh"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

type Peer struct {
	HostName    string
	ApparentIP  string
	Conn        ssh.Conn
	ID          int
	Alive       bool
	LastSeen    int64
	m           sync.Mutex
	LastAttempt int64
	NodeID      string
	MessageChan chan []byte
}

type PList struct {
	Peers     map[int]*Peer
	PeerCount int
	m         sync.Mutex
}

func (p *PList) Add(n *Peer, idoverride int) {

}

func CorrectHost(host string) string {
	return "";
}

func (p PList) ContainsIP(host string) bool {
	return false;
}

func (p PList) FindByIP(host string) int {
	return -1
}

func (p PList) RemoveByStruct(n Peer) {

}

var GlobalPeerList PList

func StartLookingForPeers() {

}

func SystemCleanup() {

}

func RestorePeerList() {

}

func ScountOutNewPeers() {

}
