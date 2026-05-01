package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nm := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])

	board := make([][]byte, n)
	for i := range n {
		scanner.Scan()
		board[i] = scanner.Bytes()
	}

	cost := make([][]int, m)
	for i := range m {
		scanner.Scan()
		fields := strings.Fields(scanner.Text())
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