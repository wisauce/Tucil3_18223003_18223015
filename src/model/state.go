package model

// import "fmt"

type State struct {
	X, Y       int
	NextNumber int
	Cost       int
	Parent     *State
}

type StateKey struct {
	X, Y       int
	NextNumber int
}

func (s State) Key() StateKey {
	return StateKey{X: s.X, Y: s.Y, NextNumber: s.NextNumber}
}
