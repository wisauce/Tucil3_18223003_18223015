package main

import (
	"IceSlidingPuzzle/utils"
	"fmt"
)

func main() {
	fmt.Println("Masukkan nama file papan permainan (.txt)")
	var filename string
	fmt.Scan(&filename)
	state, solver := utils.ParseFile(filename)
	// fmt.Print(state)
	final := solver.UCS(state)
	fmt.Print(final)
}