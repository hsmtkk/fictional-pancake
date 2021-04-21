package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/icmp"
)

func main() {
	pc, err := icmp.ListenPacket("udp4", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()
	dst, err := net.ResolveUDPAddr("192.168.11.1", "udp")
	if err != nil {
		log.Fatal(err)
	}
	_, err = pc.WriteTo([]byte("hoge"), dst)
	if err != nil {
		log.Fatal(err)
	}
	rb := make([]byte, 1500)
	_, peer, err := pc.ReadFrom(rb)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", peer)
}
