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

	final, solved := solver.UCS(state)
	if !solved {
		fmt.Println("Tidak ada solusi yang ditemukan.")
		return
	} else {
		solver.VisualizeRoute(final)
	}
}
