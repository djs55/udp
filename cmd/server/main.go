package main

import (
	"flag"
	"log"
	"net"
)

// UDPBufSize is the max UDP payload length
const UDPBufSize = 65507

func main(){
	var port int
	flag.IntVar(&port, "port", 8888, "UDP port to listen on")
	flag.Parse()

	l, err := net.ListenUDP("udp", &net.UDPAddr{
		IP: net.ParseIP("0.0.0.0"),
		Port: port,
	})
	if err != nil {
		log.Fatal("failed to bind on port %d: %v", port, err)
	}
	log.Printf("Listening on UDP port %d\n", port)
	buf := make([]byte, UDPBufSize)
	for {
		n, from, err := l.ReadFromUDP(buf)
		if err != nil {
			log.Fatal("failed to read UDP: %v", err)
		}
		log.Printf("Read '%s', replying\n", string(buf[0:n]))
		if _, err := l.WriteToUDP(buf[0:n], from); err != nil {
			log.Fatal("failed to write UDP: %v", err)
		}
	}
}