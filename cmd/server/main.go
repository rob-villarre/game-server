package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"

	c "main/internal/common"
	u "main/internal/udp"
)

var state = NewGameState()

func SendMessage(conn *net.UDPConn, addr *net.UDPAddr, msg u.Message) {
	data, _ := json.Marshal(msg)
	conn.WriteToUDP(data, addr)
}

func HandleConnect(conn *net.UDPConn, addr *net.UDPAddr) {
	id := len(state.Players) + 1

	state.Players[id] = NewPlayer(id, nil, c.Vector{X: 50, Y: 50}, 0.0)

	ack := u.Message{
		Type:      "CONNECT_ACK",
		Player:    *state.Players[id],
		Timestamp: time.Now().Unix(),
	}

	SendMessage(conn, addr, ack)
	fmt.Printf("Player %d connected from %s\n", id, addr.String())
}

func HandleClient(conn *net.UDPConn) {
	buffer := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error reading from UDP:", err)
		return
	}

	var msg u.Message
	err = json.Unmarshal(buffer[:n], &msg)
	if err != nil {
		fmt.Println("Invalid packet received:", err)
	}

	switch msg.Type {
	case "CONNECT":
		HandleConnect(conn, addr)
	}
}

func main() {

	args := os.Args
	if len(args) == 1 {
		fmt.Println("server <port>")
		return
	}

	PORT := ":" + args[1]
	udp, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	conn, err := net.ListenUDP("udp4", udp)
	if err != nil {
		fmt.Println("Error starting UDP server", err)
		return
	}
	fmt.Println("server started on port", PORT)

	defer conn.Close()

	for {
		HandleClient(conn)
	}
}
