package main

import (
	"errors"
	"fmt"
	"net"
	"os"
)

func main() {
	// Parse command line arguments
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <port>\n", os.Args[0])
		os.Exit(1)
	}
	port := os.Args[1]

	// Create UDP address for listening
	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error resolving UDP address:", err)
		os.Exit(1)
	}

	// Create UDP connection
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating UDP connection:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Wait for UDP packets
	buffer := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			// Check if the error is a timeout error and continue if it is
			var timeoutError net.Error
			if errors.As(err, &timeoutError); timeoutError.Timeout() {
				continue
			}
			fmt.Fprintln(os.Stderr, "Error receiving UDP packet:", err)
			continue
		}

		// Send a response packet with payload "pong" to the source address
		_, err = conn.WriteToUDP([]byte("pong"), clientAddr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error sending UDP packet:", err)
			continue
		}

		fmt.Printf("Received %d bytes from %s: %s\n", n, clientAddr.String(), string(buffer[:n]))
	}
}
