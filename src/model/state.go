package model

import "fmt"

type State struct {
	X, Y       int
	NextNumber int // angka berikutnya yang harus dilewati
	Cost       int // g(n)
}

func (s State) Move(direction Direction, board [][]byte, costs [][]int) (State, bool) {
	x, y := s.X, s.Y
	nextNumber := s.NextNumber
	cost := 0
	for {
		switch direction {
		case UP:
			y--
		case DOWN:
			y++
		case RIGHT:
			x++
		case LEFT:
			x--
		}
		if y <= 0 || x <= -1 || y >= len(board) || x >= len(board[y]) || board[y][x] == 'L' || board[y][x] == 'O' || board[y][x] == 'X' {
			break
		}
		fmt.Println(x, y)

		cost += costs[y][x]
    if board[y][x] == byte(nextNumber+'0') {
			nextNumber++
			board[y][x] = '*'
		}
	}
	if y >= len(board) || x >= len(board[y]) || board[y][x] == 'L' {
		return State{}, false
	}

  if board[y][x] == 'X' {
    switch direction {
		case UP:
			y++
		case DOWN:
			y--
		case RIGHT:
			x--
		case LEFT:
			x++
		}
  }

	return State{
		X:          x,
		Y:          y,
		NextNumber: nextNumber,
		Cost:       cost,
	}, true
}