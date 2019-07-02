package main

import (
	"flag"
	"log"
	"net"
	"time"
)

// UDPBufSize is the max UDP payload length
const UDPBufSize = 65507

func main(){
	var address, payload string
	flag.StringVar(&address, "address", "localhost:8888", "address to send UDP to")
	flag.StringVar(&payload, "payload", "this is a test", "payload message to send")
	flag.Parse()

	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatalf("unable to parse address %s: %v", address, err)
	}
	c, err := net.DialUDP("udp",nil, addr)
	if err != nil {
		log.Fatal("failed to Dial to UDP %s: %v", address, err)
	}

	msg := []byte(payload)
	buf := make([]byte, UDPBufSize)

	for {
		n, err := c.Write(msg)
		if err != nil {
			log.Fatal("failed to send UDP: %v", err)
		}
		if err := c.SetReadDeadline(time.Now().Add(time.Second)); err != nil {
			log.Fatal("failed to set read deadline: %v", err)
		}
		n, err = c.Read(buf)
		if err != nil {
			log.Println("missed UDP response")
			continue
		}
		received := string(buf[0:n])
		log.Printf("Read '%s'\n", received)
		if received != payload {
			log.Fatalf("Received unexpected response. I sent '%s' and received '%s'", payload, received)
		}
		time.Sleep(time.Second)
	}
}
