package main

import (
	c "main/internal/common"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Id           int
	Position     c.Vector
	Speed        float64
	Velocity     c.Vector
	TurnRate     float64
	Thrust       float64
	Heading      float64
	Decay        float64
	Sprite       *ebiten.Image
	EngineEffect *ebiten.Image

	ThrustRange   [2]float64 // Min and Max thrust (m/s^2)
	TurnRateRange [2]float64 // Min and Max turn rate (degrees/sec)
	MaxSpeed      float64    // Max speed (m/s)
	Drag          float64    // Drag coefficient (m/s^2)
	Radius        float64    // Object radius (meters)
	Mass          float64    // Object mass (kg)
}

func NewPlayer(id int, position c.Vector, angle float64, sprite *ebiten.Image, engineEffect *ebiten.Image) *Player {
	return &Player{
		Id:           id,
		Position:     position,
		Speed:        0.0,
		TurnRate:     0.0,
		Velocity:     c.Vector{X: 0.0, Y: 0.0},
		Heading:      angle,
		Sprite:       sprite,
		EngineEffect: engineEffect,

		ThrustRange:   [2]float64{-480.0, 480.0},
		TurnRateRange: [2]float64{-180.0, 180.0},
		MaxSpeed:      240.0,
		Drag:          80.0,
		Radius:        20.0,
		Mass:          300.0,
	}
}

func Sign(x float64) float64 {
	if x > 0 {
		return 1.0
	} else if x < 0 {
		return -1.0
	}
	return 0.0
}

func (p *Player) Update() error {

	// dt := 1.0 / ebiten.ActualFPS()
	dt := 1.0 / 60.0
	p.TurnRate = 0.0
	p.Thrust = 0.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.TurnRate = -180.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.TurnRate = 180.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Thrust = 150.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Thrust = -150.0
	}

	drag_amount := p.Drag * dt
	if drag_amount > math.Abs(p.Speed) {
		p.Speed = 0.0
	} else {
		p.Speed -= drag_amount * Sign(p.Speed)
	}

	// Bounds check the thrust
	if p.Thrust < p.ThrustRange[0] || p.Thrust > p.ThrustRange[1] {
		p.Thrust = math.Min(math.Max(p.ThrustRange[0], p.Thrust), p.ThrustRange[1])
	}

	// Apply thrust
	p.Speed += p.Thrust * dt

	// Bounds check the speed
	if p.Speed > p.MaxSpeed {
		p.Speed = p.MaxSpeed
	} else if p.Speed < -p.MaxSpeed {
		p.Speed = -p.MaxSpeed
	}

	// Bounds check the turn rate
	if p.TurnRate < p.TurnRateRange[0] || p.TurnRate > p.TurnRateRange[1] {
		p.TurnRate = math.Min(math.Max(p.TurnRateRange[0], p.TurnRate), p.TurnRateRange[1])
	}

	// Update the angle based on turning rate
	p.Heading += p.TurnRate * dt

	p.Heading = math.Mod(p.Heading, 360.0)
	if p.Heading < 0 {
		p.Heading += 360.0
	}

	// Use speed magnitude to get velocity vector
	headingRad := p.Heading * (math.Pi / 180)
	p.Velocity = c.Vector{
		X: math.Sin(headingRad) * p.Speed,
		Y: -(math.Cos(headingRad) * p.Speed),
	}

	// Update the position based off the velocities
	p.Position = c.Vector{
		X: p.Position.X + p.Velocity.X*dt,
		Y: p.Position.Y + p.Velocity.Y*dt,
	}

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	bounds := p.Sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	// Scale before translating to avoid weird offsets
	op.GeoM.Scale(0.5, 0.5)

	// Move the origin to the center for proper rotation
	op.GeoM.Translate(-halfW*0.5, -halfH*0.5)

	headingRad := p.Heading * (math.Pi / 180)
	op.GeoM.Rotate(headingRad)

	// Translate to actual screen position
	op.GeoM.Translate(p.Position.X, p.Position.Y)

	if p.Thrust != 0.0 {
		engineOp := &ebiten.DrawImageOptions{}
		engineOp.GeoM.Scale(0.5, 0.5)
		engineOp.GeoM.Translate(-halfW*0.5, -halfH*0.5)
		engineOp.GeoM.Rotate(headingRad)

		// Offset engine effect behind the ship
		engineOffset := float64(bounds.Dy())*0.5 - 10.0
		engineX := p.Position.X - math.Sin(headingRad)*engineOffset
		engineY := p.Position.Y + math.Cos(headingRad)*engineOffset

		engineOp.GeoM.Translate(engineX, engineY)

		screen.DrawImage(p.EngineEffect, engineOp)
	}

	screen.DrawImage(p.Sprite, op)
}
