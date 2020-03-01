package main

import (
	"fmt"
	"net"
	"os"

	iptables "github.com/coreos/go-iptables/iptables"
)

const ProtocolIPv4 Protocol = iota

func main() {
	ipt, err := iptables.New()
	if err != nil {
		fmt.Println("New failed: ")
	}

	fmt.Println()
}
