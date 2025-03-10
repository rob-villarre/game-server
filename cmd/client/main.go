package main

import (
	"encoding/json"
	"fmt"
	"log"
	"main/cmd/client/assets"
	"net"
	"os"
	"strconv"

	c "main/internal/common"
	u "main/internal/udp"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

var conn *net.UDPConn

type Game struct {
	player *Player
}

func (g *Game) Update() error {

	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "FPS: "+strconv.FormatFloat(ebiten.ActualFPS(), 'f', 1, 64))

	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func connect() *net.UDPConn {
	arguments := os.Args
	if len(arguments) == 1 {
		log.Fatalln("Please provide a host:port string")
	}

	CONNECT := arguments[1]

	udp, err := net.ResolveUDPAddr("udp4", CONNECT)
	if err != nil {
		log.Fatalln("Error resolving UDP address:", err)
	}

	conn, err := net.DialUDP("udp4", nil, udp)
	if err != nil {
		log.Fatalln("Error dialing udp address", err)
	}

	fmt.Printf("The UDP server is %s\n", conn.RemoteAddr().String())

	return conn
}

func main() {

	conn := connect()
	defer conn.Close()

	// TODO: add authentication
	msg := u.Message{
		Type: "CONNECT",
	}
	data, _ := json.Marshal(msg)
	_, err := conn.Write(data)
	if err != nil {
		log.Fatalln(err)
	}

	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	var response u.Message
	json.Unmarshal([]byte(buffer[0:n]), &response)
	playerBytes, _ := json.Marshal(response.Player)
	var player Player
	json.Unmarshal(playerBytes, &player)

	fmt.Println(player)

	g := &Game{
		player: NewPlayer(
			player.Id,
			c.Vector{X: player.Position.X, Y: player.Position.Y},
			player.Heading,
			assets.PlayerSprite,
			assets.PlayerEngineEffect,
		),
	}

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
