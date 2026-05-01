package main

import (
	"IceSlidingPuzzle/utils"
	"fmt"
)

func main() {
	fmt.Println("Masukkan nama file papan permainan (.txt)")
	var filename string
	fmt.Scan(&filename)
	board, cost, state := utils.ParseFile(filename)

	for i := range board {
		fmt.Println(string(board[i]))
	}
	for i := range cost {
		fmt.Println(cost[i])
	}
	fmt.Print(state)
}