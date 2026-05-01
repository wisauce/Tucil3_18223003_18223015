package main

import (
	"IceSlidingPuzzle/model"
	"IceSlidingPuzzle/utils"
	"fmt"
)

func main() {
	fmt.Println("Masukkan nama file papan permainan (.txt)")
	var filename string
	fmt.Scan(&filename)
	state, board, costs := utils.ParseFile(filename)
	for i := range board {
		fmt.Println(string(board[i]))
	}
	for i := range costs {
		fmt.Println(costs[i])
	}
	fmt.Print(state)
	newstate, canMove := state.Move(model.UP, board, costs)
	if canMove {
		fmt.Print(newstate)
	}
}