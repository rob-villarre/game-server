package main

import (
	"fmt"
	"math/rand/v2"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("server <port>")
		return
	}

	address := ":" + args[1]
	udp_address, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.ListenUDP("udp4", udp_address)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
	buffer := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("-> ", string(buffer[0:n-1]))

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return
		}

		r := rand.IntN(100-1) + 1
		data := []byte(strconv.Itoa(r))
		fmt.Printf("data: %s\n", string(data))
		_, err = conn.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	// fmt.Printf("The UDP server is %s\n", conn.RemoteAddr().String())
}
