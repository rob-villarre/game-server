package main

import (
	"log"
	"main/cmd/client/assets"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Vector struct {
	X float64
	Y float64
}

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

type Message struct {
	Type      string  `json:"type"`
	PlayerId  int     `json:"player_id,omitempty"`
	X         float64 `json:"x,omitempty"`
	Y         float64 `json:"y,omitempty"`
	Timestamp int64   `json:"timestamp,omitempty"`
}

func main() {

	g := &Game{
		player: NewPlayer(0, Vector{100, 100}, 0.0, assets.PlayerSprite),
	}

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

	// arguments := os.Args
	// if len(arguments) == 1 {
	// 	fmt.Println("Please provide a host:port string")
	// 	return
	// }
	// CONNECT := arguments[1]

	// s, err := net.ResolveUDPAddr("udp4", CONNECT)
	// c, err := net.DialUDP("udp4", nil, s)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	// defer c.Close()

	// msg := Message{
	// 	Type: "CONNECT",
	// }
	// data, _ := json.Marshal(msg)
	// _, err = c.Write(data)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// buffer := make([]byte, 1024)
	// n, _, err := c.ReadFromUDP(buffer)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("Reply: %s\n", string(buffer[0:n]))
}
