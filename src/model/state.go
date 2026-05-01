package model

// import "fmt"

type State struct {
	X, Y       int
	NextNumber int
	Cost       int
  Parent *State
}

type StateKey struct {
  X, Y int
  NextNumber int
}