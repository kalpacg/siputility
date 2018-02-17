package siputility

import (
	"fmt"
	"log"
	"net"
	"strings"
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

// Returns MTU of the given network interface.
func MTU(ifAddr string) int {
	ifAddrs, err := net.Interfaces()

	if err != nil {
		log.Fatalf("Error in figuring out interfaces for system", err)
	}

	for _, i := range ifAddrs {
		if i.Name == ifAddr {
			return i.MTU
		}
	}
	return 1024
}

// Returns IP address assigned to a given network interface
func AssignedIP(ifAddr string) net.IP {
	ifAddrs, err := net.Interfaces()

	if err != nil {
		log.Fatalf("Error in figuring out interfaces for system", err)
	}

	for _, i := range ifAddrs {
		if i.Name == ifAddr {
			// Get all unicast addresses for i
			addrs, err := i.Addrs()

			if err != nil {
				log.Fatalf("Error in figuring out address for interface", err)
			}

			for _, addr := range addrs {
				ip, _, err := net.ParseCIDR(addr.String())

				if err != nil {
					log.Fatalf("Error in figuring out address for interface", err)
				}

				if IsIPv4(&ip) {
					return ip
				}
			}
		}
	}
	return nil
}

// If IP address is ipv4 type, returns true. Otherwise false.
func IsIPv4(ip *net.IP) bool {
	return strings.Contains(ip.String(), ".")
}
