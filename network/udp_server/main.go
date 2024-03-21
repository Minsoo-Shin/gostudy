package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func MakeUDPServer() net.Addr {
	server, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("server address: ", server.LocalAddr().String())
	go func() {
		for {
			buf := make([]byte, 1024)
			n, clientAddress, err := server.ReadFrom(buf)
			if n == 0 {
				fmt.Println("buffer size is 0...")
			}
			if n > 0 {
				fmt.Println("server received: ", string(buf[:n]), "from", clientAddress)
			}
			if err != nil {
				log.Fatal(err)
			}
			_, err = server.WriteTo(buf[:n], clientAddress)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()
	return server.LocalAddr()
}
func main() {
	serverAddr := MakeUDPServer()
	client, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("client address: ", client.LocalAddr())
	time.Sleep(3 * time.Second)
	client.WriteTo([]byte("ping"), serverAddr)
	buf := make([]byte, 1024)
	n, addr, _ := client.ReadFrom(buf)
	fmt.Println("client received: ", string(buf[:n]), addr)
}
