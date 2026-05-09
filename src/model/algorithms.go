package model

import (
	"container/heap"
)

type AlgorithmResult struct {
	FinalState State
	Solved     bool
	Iterations int
}

type HeuristicFunc func(st State, s Solver) int

func (s Solver) UCS(st State) AlgorithmResult {
	return s.searchWithPriorityQueue(st, func(state State) int {
		return state.Cost
	})
}

func (s Solver) GBFS(st State, h HeuristicFunc) AlgorithmResult {
	return s.searchWithPriorityQueue(st, func(state State) int {
		return h(state, s)
	})
}

func (s Solver) AStar(st State, h HeuristicFunc) AlgorithmResult {
	return s.searchWithPriorityQueue(st, func(state State) int {
		return state.Cost + h(state, s)
	})
}

func (s Solver) searchWithPriorityQueue(st State, calculatePriority func(State) int) AlgorithmResult {
	open := &PriorityQueue{}
	heap.Init(open)

	heap.Push(open, &Item{
		State:    &st,
		Priority: calculatePriority(st),
	})

	close := make(map[StateKey]int)
	iterations := 0

	for open.Len() > 0 {
		iterations++
		item := heap.Pop(open).(*Item)
		cur := item.State

		key := cur.Key()

		if oldCost, ok := close[key]; ok && oldCost <= cur.Cost {
			continue
		}
		close[key] = cur.Cost

		if s.isSolved(*cur) {
			return AlgorithmResult{FinalState: *cur, Solved: true, Iterations: iterations}
		}

		nextStates := s.generateNextmoves(*cur)

		for i := range nextStates {
			ns := nextStates[i]
			heap.Push(open, &Item{
				State:    &ns,
				Priority: calculatePriority(ns),
			})
		}
	}

	return AlgorithmResult{Solved: false, Iterations: iterations}
}

func (s Solver) BFS(st State) AlgorithmResult {
	queue := []State{st}
	visited := make(map[StateKey]bool)
	iterations := 0

	for len(queue) > 0 {
		iterations++
		cur := queue[0]
		queue = queue[1:]

		key := cur.Key()
		if visited[key] {
			continue
		}
		visited[key] = true

		if s.isSolved(cur) {
			return AlgorithmResult{FinalState: cur, Solved: true, Iterations: iterations}
		}

		nextStates := s.generateNextmoves(cur)
		for _, ns := range nextStates {
			if !visited[ns.Key()] {
				queue = append(queue, ns)
			}
		}
	}

	return AlgorithmResult{Solved: false, Iterations: iterations}
}
