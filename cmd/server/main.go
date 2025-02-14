package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

type Message struct {
	Type      string  `json:"type"`
	PlayerId  int     `json:"player_id,omitempty"`
	X         float64 `json:"x,omitempty"`
	Y         float64 `json:"y,omitempty"`
	Timestamp int64   `json:"timestamp,omitempty"`
}

type Player struct {
	Id   int
	Addr *net.UDPAddr
	X, Y float64
}

var players = make(map[int]*Player)

func SendMessage(conn *net.UDPConn, addr *net.UDPAddr, msg Message) {
	data, _ := json.Marshal(msg)
	conn.WriteToUDP(data, addr)
}

func HandleConnect(conn *net.UDPConn, addr *net.UDPAddr) {
	id := len(players) + 1

	players[id] = &Player{
		Id:   id,
		Addr: addr,
		X:    0,
		Y:    0,
	}

	ack := Message{
		Type:      "CONNECT_ACK",
		PlayerId:  id,
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

	var msg Message
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
