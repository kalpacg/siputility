package siputility

import (
	"fmt"
	"log"
	"net"
)

func StartConnection(network, address string) *net.UDPConn {
	// Try to return an UDP endpoint
	udpAddr, err := net.ResolveUDPAddr(network, address)

	if err != nil {
		log.Fatalf("Could not resolve an UDP endpoint", err)
	}

	udpConn, err := net.ListenUDP(network, udpAddr)

	if err != nil {
		log.Fatalf("Could not create an UDP listening port", err)
	}

	return udpConn
}

func ReadData(udpConn *net.UDPConn, buffer []byte) {

	recvBytes, _, err := udpConn.ReadFromUDP(buffer)

	if err != nil {
		log.Fatalf("Data read error from UDP connection", err)
	}

	fmt.Println(string(buffer[0:recvBytes]))
}
