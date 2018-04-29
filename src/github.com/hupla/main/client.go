package main

import (
	"github.com/hupla/conf"
	"log"
	"net"
)

//exercise multiple ways of sending a broadcast request
//used for testing server.go broadcast listener
func main() {

	message := []byte("hello")
	ipBcast := net.IPv4(192, 168, 1, 255)
	//ipBcast := net.IPv4(255, 255, 255, 255)

	//set local ip binding
	ipLocal := net.IPv4(192, 168, 1, 10)
	//ipLocal := net.IPv4(10, 10, 10, 129)
	//ipLocal := net.IPv4(127, 0, 0, 1)
	lAddr := net.UDPAddr{IP: ipLocal, Port: 8002}

	//if you only have one NIC you can just leave &lAddr as nil
	socket, err := net.DialUDP("udp4", &lAddr, &net.UDPAddr{
		IP:   ipBcast,
		Port: conf.BPort,
	})
	if err != nil {
		log.Printf("Error while starting TXQueue: %s, \n%v\n %v\n", err, ipBcast, ipLocal)
		return
	}

	socket.Write(message)
}
