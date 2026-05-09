package model

import "math"

func FindTargets(s Solver) map[int][2]int {
	targets := make(map[int][2]int)
	for y := 0; y < len(s.Board); y++ {
		for x := 0; x < len(s.Board[0]); x++ {
			c := s.Board[y][x]
			if c >= '0' && c <= '9' {
				targets[int(c-'0')] = [2]int{x, y}
			}
		}
	}
	targets[s.FinalNumber+1] = [2]int{s.GoalX, s.GoalY}
	return targets
}

func H1Manhattan(targets map[int][2]int) HeuristicFunc {
	return func(st State, s Solver) int {
		return calculateDistance(st, s, targets, func(dx, dy float64) float64 {
			return math.Abs(dx) + math.Abs(dy)
		})
	}
}

func H2Euclidean(targets map[int][2]int) HeuristicFunc {
	return func(st State, s Solver) int {
		return calculateDistance(st, s, targets, func(dx, dy float64) float64 {
			return math.Sqrt(dx*dx + dy*dy)
		})
	}
}

func H3Chebyshev(targets map[int][2]int) HeuristicFunc {
	return func(st State, s Solver) int {
		return calculateDistance(st, s, targets, func(dx, dy float64) float64 {
			return math.Max(math.Abs(dx), math.Abs(dy))
		})
	}
}

func H4MissingTargets() HeuristicFunc {
	return func(st State, s Solver) int {
		return (s.FinalNumber + 1 - st.NextNumber)
	}
}

func H5StraightToGoal(targets map[int][2]int) HeuristicFunc {
	return func(st State, s Solver) int {
		goalTarget := targets[s.FinalNumber+1]
		dx := math.Abs(float64(st.X - goalTarget[0]))
		dy := math.Abs(float64(st.Y - goalTarget[1]))
		return int(dx + dy)
	}
}

func calculateDistance(st State, s Solver, targets map[int][2]int, distFunc func(dx, dy float64) float64) int {
	cost := 0.0
	currX, currY := st.X, st.Y
	for i := st.NextNumber; i <= s.FinalNumber+1; i++ {
		target, exists := targets[i]
		if !exists {
			continue
		}
		dx := float64(currX - target[0])
		dy := float64(currY - target[1])
		cost += distFunc(dx, dy)
		currX, currY = target[0], target[1]
	}
	return int(cost)
}
