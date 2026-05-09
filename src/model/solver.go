package model

import (
	"strings"
)

type Solver struct {
	Board        [][]byte
	Costs        [][]int
	GoalX, GoalY int
	FinalNumber  int
}



func (s Solver) isSolved(st State) bool {
	return st.X == s.GoalX && st.Y == s.GoalY && st.NextNumber == s.FinalNumber+1
}

func (s Solver) generateNextmoves(st State) []State {
	moves := make([]State, 0, 4)
	for _, d := range []rune{'U', 'D', 'L', 'R'} {
		state, canmove := s.Move(d, st)
		if canmove {
			moves = append(moves, state)
		}
	}
	return moves
}
func (s Solver) Move(dir rune, st State) (State, bool) {
    x, y := st.X, st.Y
    nextNumber := st.NextNumber
    cost := st.Cost

    for {
        nx, ny := x, y
        switch dir {
			case 'U':  ny--
			case 'D':  ny++
			case 'L':  nx--
			case 'R':  nx++
        }

        if ny < 0 || nx < 0 || ny >= len(s.Board) || nx >= len(s.Board[0]) || s.Board[ny][nx] == 'L' {
            return State{}, false
        }
        if s.Board[ny][nx] == 'X' {
            break
        }

        x, y = nx, ny
        cost += s.Costs[y][x]  

        if s.Board[y][x] >= '0' && s.Board[y][x] <= '9' {
            currentNum := int(s.Board[y][x] - '0')
        	if currentNum == nextNumber {
                nextNumber++
            } else if currentNum > nextNumber {
                return State{}, false
            }
        }
    }

    if st.X == x && st.Y == y {
        return State{}, false
    }

    return State{X: x, Y: y, NextNumber: nextNumber, Cost: cost, Parent: &st}, true
}

func (s Solver) VisualizeState(st State) string {
	var sb strings.Builder
	for i := range len(s.Board) {
		for j := range len(s.Board[0]) {
			if j == st.X && i == st.Y {
				sb.WriteByte('Z')
				continue
			}

			tile := s.Board[i][j]
			if tile >= '0' && tile <= '9' {
				num := int(tile - '0')
				if num < st.NextNumber {
					sb.WriteByte('*')
					continue
				}
			}

			sb.WriteByte(tile)
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (s Solver) GetRoute(st State) ([]*State, []rune) {
	states := []*State{}
	moves := []rune{}

	current := &st
	for current != nil {
		states = append(states, current)
		if current.Parent != nil {
			dx := current.X - current.Parent.X
			dy := current.Y - current.Parent.Y

			var dir rune
			if dx > 0 {
				dir = 'R'
			} else if dx < 0 {
				dir = 'L'
			} else if dy < 0 {
				dir = 'U'
			} else if dy > 0 {
				dir = 'D'
			} else {
				dir = '?'
			}
			moves = append(moves, dir)
		}
		current = current.Parent
	}

	for i, j := 0, len(states)-1; i < j; i, j = i+1, j-1 {
		states[i], states[j] = states[j], states[i]
	}
	for i, j := 0, len(moves)-1; i < j; i, j = i+1, j-1 {
		moves[i], moves[j] = moves[j], moves[i]
	}

	return states, moves
}
