package main

import (
	c "main/internal/common"
	"net"
)

type Player struct {
	Id       int
	Conn     *net.UDPConn
	Position c.Vector
	Speed    float64
	Velocity c.Vector
	TurnRate float64
	Thrust   float64
	Heading  float64
	Decay    float64

	ThrustRange   [2]float64 // Min and Max thrust (m/s^2)
	TurnRateRange [2]float64 // Min and Max turn rate (degrees/sec)
	MaxSpeed      float64    // Max speed (m/s)
	Drag          float64    // Drag coefficient (m/s^2)
	Radius        float64    // Object radius (meters)
	Mass          float64    // Object mass (kg)
}

func NewPlayer(id int, conn *net.UDPConn, position c.Vector, angle float64) *Player {
	return &Player{
		Id:       id,
		Conn:     conn,
		Position: position,
		Speed:    0.0,
		TurnRate: 0.0,
		Velocity: c.Vector{X: 0.0, Y: 0.0},
		Heading:  angle,

		ThrustRange:   [2]float64{-480.0, 480.0},
		TurnRateRange: [2]float64{-180.0, 180.0},
		MaxSpeed:      240.0,
		Drag:          80.0,
		Radius:        20.0,
		Mass:          300.0,
	}
}
