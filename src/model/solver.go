package model

import (
	"container/heap"
	"fmt"
)

type Solver struct {
	Board        [][]byte
	Costs        [][]int
	GoalX, GoalY int
	FinalNumber  int
}

func (s Solver) UCS(st State) State {
	open := &PriorityQueue{}
	heap.Init(open)

	heap.Push(open, &Item{
		State:    &st,
		Priority: st.Cost,
	})

	close := make(map[StateKey]int)

	for open.Len() > 0 {
		item := heap.Pop(open).(*Item)
		cur := item.State

		key := StateKey{cur.X, cur.Y, cur.NextNumber}

		if oldCost, ok := close[key]; ok && oldCost <= cur.Cost {
			continue
		}

		close[key] = cur.Cost

		if s.isSolved(*cur) {
			return *cur
		}

		nextStates := s.generateNextmoves(*cur)

		for i := range nextStates {
			ns := nextStates[i]

			nextKey := StateKey{ns.X, ns.Y, ns.NextNumber}

			if oldCost, ok := close[nextKey]; ok && oldCost <= ns.Cost {
				continue
			}

			heap.Push(open, &Item{
				State:    &ns,
				Priority: ns.Cost,
			})
		}
	}

	return State{}
}

func (s Solver) isSolved(st State) bool {
	return st.X == s.GoalX && st.Y == s.GoalY && st.NextNumber == s.FinalNumber+1
}

func (s Solver) generateNextmoves(st State) []State {
	moves := make([]State, 0, 4)
	for i := range 4 {
		state, canmove := s.move(Direction(i), st)
		if canmove {
			moves = append(moves, state)
		}
	}
	return moves
}

func (s Solver) move(dir Direction, st State) (State, bool) {
	x, y := st.X, st.Y
	nextNumber := st.NextNumber
	cost := st.Cost
	for {
		cost += s.Costs[y][x]

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
		if s.Board[y][x] >= '0' && s.Board[y][x] <= '9' {
			if s.Board[y][x] == byte(nextNumber+'0') {
				nextNumber++
			} else {
				return State{}, false
			}
		}
	}
	if st.X == x && st.Y == y {
		return State{}, false
	}

	return State{
		X:          x,
		Y:          y,
		NextNumber: nextNumber,
		Cost:       cost,
		Parent:     &st,
	}, true
}

func (s Solver) VisualizeState(st State) {
	for i := range len(s.Board) {
		for j := range len(s.Board[0]) {
			if j == st.X && i == st.Y {
				fmt.Print("Z")
				continue
			}

			tile := s.Board[i][j]
			if tile >= '0' && tile <= '9' {
				num := int(tile - '0')
				if num < st.NextNumber {
					fmt.Print("*")
					continue
				}
			}

			fmt.Printf("%c", tile)
		}
		fmt.Println()
	}
}

func (s Solver) VisualizeRoute(st State) {
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

	fmt.Printf("Solusi Yang Ditemukan : %s\n", string(moves))
	fmt.Printf("Cost dari Solusi : %d\n", st.Cost)
	fmt.Println("Initial")
	s.VisualizeState(*states[0])
	fmt.Println()

	for i := 1; i < len(states); i++ {
		fmt.Printf("Step %d : %c\n", i, moves[i-1])
		s.VisualizeState(*states[i])
		fmt.Println()
	}
}
