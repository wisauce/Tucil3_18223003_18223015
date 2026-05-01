package model

import (
	"container/heap"
)

type Solver struct {
  Board [][]byte
  Costs [][]int
  GoalX, GoalY int
}

func (sol Solver) isSolved(s State) bool {
    return s.X == sol.GoalX && s.Y == sol.GoalY
}

func (s Solver) UCS (st State) State {
	open := &PriorityQueue{}
	heap.Init(open)
	heap.Push(open, &Item{
		State:    &st,
		Priority: st.Cost,
	})
	close := make(map[StateKey]int)

	item := heap.Pop(open).(*Item)
	cur := item.State

	for !s.isSolved(*cur) {
		key := StateKey{
			X: cur.X,
			Y: cur.Y,
			NextNumber: cur.NextNumber,
		}
		close[key] = cur.Cost

		nextStates := s.generateNextmoves(*cur)
		for _, state := range nextStates {
			heap.Push(open, &Item{
				State: &state,
				Priority: state.Cost,
			})
		}

		item = heap.Pop(open).(*Item)
		cur = item.State
	}

	return *cur
}

func (s Solver) generateNextmoves(st State) []State {
	moves := make([]State,0,4)
	for i := range 4 {
		state, canmove := s.move(Direction(i),st)
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
		Parent: &st,
	}, true
}