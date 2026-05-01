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
	state, solver := utils.ParseFile(filename)
	fmt.Print(state)
	newstate, canMove := solver.Move(model.UP, state)
	if canMove {
		fmt.Print(newstate)
	}
}