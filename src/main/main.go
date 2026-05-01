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

	final := solver.UCS(state)
	solver.VisualizeRoute(final)
}
