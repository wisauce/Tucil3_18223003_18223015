package utils

import (
	"IceSlidingPuzzle/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseFile(filename string) (model.State, model.Solver) {
	file, err := os.Open(("../test/" + filename))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	nm := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])

	var initial model.State
	var solver model.Solver

	board := make([][]byte, n)
	numberSequence := [10]bool{}
	for i := range n {
		scanner.Scan()
		line := scanner.Text()
		if len(line) != m {
			panic("invalid board input width")
		}
		board[i] = make([]byte, m)
		for j := range m {
			c := line[j]
			if c >= '0' && c <= '9' {
				if numberSequence[c-'0'] {
					panic("overlapping <i> detected")
				}
				numberSequence[c-'0'] = true
			} else if c != '*' && c != 'X' && c != 'L' && c != 'Z' && c != 'O' {
				panic(fmt.Sprintf("%c is not a valid tile", c))
			}
			switch c {
			case 'Z':
				initial.X = j
				initial.Y = i
				board[i][j] = '*'
			case 'O':
				solver.GoalX = j
				solver.GoalY = i
				board[i][j] = 'O'
			default:
				board[i][j] = c
			}
		}
	}
	changed := false
	prevNumber := numberSequence[0]
	if prevNumber {
		initial.NextNumber = 0
	}
	for i := 1; i < len(numberSequence); i++ {
		if numberSequence[i] != prevNumber {
			if !changed {
				if !prevNumber {
					panic("given <i> is not sequencial")
				}
				changed = true
				solver.FinalNumber = i-1
			} else if changed {
				panic("given <i> is not sequencial")
			}
		}
		prevNumber = numberSequence[i]
	}

	costs := make([][]int, n)
	for i := range n {
		scanner.Scan()
		fields := strings.Fields(scanner.Text())
		if len(fields) != m {
			panic("invalid costs input width")
		}
		costs[i] = make([]int, m)
		for j := range m {
			integer, _ := strconv.Atoi(fields[j])
			costs[i][j] = integer
		}
	}
	solver.Board = board
	solver.Costs = costs
	return initial, solver
}
