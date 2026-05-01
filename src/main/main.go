package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Masukkan nama file papan permainan (.txt)")
	var filename string
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	nm := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])

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
				if numberSequence[c - '0'] {
					panic("overlapping <i> detected")
				}
				numberSequence[c - '0'] = true
			} else if c != '*' && c != 'X' && c != 'L' && c != 'Z' && c != 'O' {
				panic(fmt.Sprintf("%c is not a valid tile", c))
			}
			board[i][j] = line[j]
		}
	}
	changed := false
	prevNumber := numberSequence[0]
	for i := 1; i < len(numberSequence); i++ {
		if numberSequence[i] != prevNumber {
			if !changed {
				if !prevNumber {
					panic("given <i> is not sequencial")
				}
				changed = true
			} else if changed {
				panic("given <i> is not sequencial")
			}
		}
		prevNumber = numberSequence[i]
	}

	cost := make([][]int, n)
	for i := range n {
		scanner.Scan()
		fields := strings.Fields(scanner.Text())
		if len(fields) != m {
			panic("invalid cost input width")
		}
		cost[i] = make([]int, m)
		for j := range m {
			integer, _ := strconv.Atoi(fields[j])
			cost[i][j] = integer
		}
	}

	for i := range board {
		fmt.Println(string(board[i]))
	}
	for i := range cost {
		fmt.Println(cost[i])
	}
}