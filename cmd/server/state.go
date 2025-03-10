package main

type GameState struct {
	Players map[int]*Player
}

func NewGameState() *GameState {
	return &GameState{
		Players: make(map[int]*Player),
	}
}
