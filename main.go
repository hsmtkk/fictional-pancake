package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/ipv4"
	"golang.org/x/net/icmp"
)

func main() {
	pc, err := icmp.ListenPacket("udp4", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()
	dst, err := net.ResolveUDPAddr("udp", "192.168.11.1:12345")
	if err != nil {
		log.Fatal(err)
	}
	body := icmp.Echo {
		ID: 0,
		Seq: 0,
		Data: []byte("hoge"),
	}
	msg := icmp.Message {
		Type: ipv4.ICMPTypeEcho              ,
		Code:0,
		Body: &body,
	}
	msgBytes, err := msg.Marshal(nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = pc.WriteTo(msgBytes, dst)
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
