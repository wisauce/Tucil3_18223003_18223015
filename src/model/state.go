package model

// import "fmt"

type State struct {
	X, Y       int
	NextNumber int
	Cost       int
}

func (s State) Move(dir Direction, board [][]byte, costs [][]int) (State, bool) {
	x, y := s.X, s.Y
	nextNumber := s.NextNumber
	cost := 0
	for {
    nx, ny := x, y
    switch dir {
    case UP:
        ny--
    case DOWN:
        ny++
    case LEFT:
        nx--
    case RIGHT:
        nx++
    }
		if ny < 0 || nx < 0 || ny >= len(board) || nx >= len(board[0]) || board[ny][nx] == 'L' {
			return State{}, false
		}

    if board[ny][nx] == 'X' {
      break
    }

    x, y = nx, ny

		cost += costs[y][x]
    if board[y][x] == byte(nextNumber+'0') {
			nextNumber++
			board[y][x] = '*'
		}
	}

	return State{
		X:          x,
		Y:          y,
		NextNumber: nextNumber,
		Cost:       cost,
	}, true
}