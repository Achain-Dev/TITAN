package main

import (
	"code.google.com/p/go.crypto/ssh"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

var UnlockTimeout int64

// Returns true if the system is in super user unlock mode.
func GetUnlockedState() bool {
	return nil
}

func ProcessRPCPacket(inbound PeerPacket) {
}
