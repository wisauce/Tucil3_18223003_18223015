package model

type Solver struct {
  Board [][]byte
  Costs [][]int
  GoalX int
  GoalY int
}

func (sol Solver) IsSolved(s State) bool {
    return s.X == sol.GoalX && s.Y == sol.GoalY
}

func (s Solver) Move(dir Direction, st State) (State, bool) {
	x, y := st.X, st.Y
	nextNumber := st.NextNumber
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
		if ny < 0 || nx < 0 || ny >= len(s.Board) || nx >= len(s.Board[0]) || s.Board[ny][nx] == 'L' {
			return State{}, false
		}
    if s.Board[ny][nx] == 'X' {
      break
    }

    x, y = nx, ny
		cost += s.Costs[y][x]
    if s.Board[y][x] == byte(nextNumber+'0') {
			nextNumber++
			s.Board[y][x] = '*'
		}
	}

	return State{
		X:          x,
		Y:          y,
		NextNumber: nextNumber,
		Cost:       cost,
	}, true
}