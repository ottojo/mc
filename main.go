package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	l, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
	}

	for {
		b := make([]byte, 1024)
		n, _, err := l.ReadFromUDP(b)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(hex.Dump(b[:n]))
	}
}
