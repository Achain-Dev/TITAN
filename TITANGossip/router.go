package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

var GlobalResvChan chan []byte
var PacketCache map[int]*PeerPacket

func RelayPackets() {
}

func PacketRouter(inbound chan PeerPacket) {
}

func SendPacket(P PeerPacket) {
	
}

// Check to see if the packet has been already seen,
// if it has not then it will add it to the cache (evicting old stuff if needed)
func SeenPacketBefore(P PeerPacket) bool {
	return false
}
