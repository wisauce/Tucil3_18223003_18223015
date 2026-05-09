package main

import (
	"IceSlidingPuzzle/model"
	"IceSlidingPuzzle/utils"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	fmt.Println(">> Masukkan nama file papan permainan (.txt):")
	var filename string
	fmt.Scan(&filename)

	state, solver := utils.ParseFile(filename)

	fmt.Println(">> Algoritma apa yang anda pilih? (UCS/GBFS/A*/BFS)")
	var algo string
	fmt.Scan(&algo)
	algo = strings.ToUpper(algo)

	var heuristic model.HeuristicFunc
	if algo == "GBFS" || algo == "A*" {
		fmt.Println(">> Heuristic apa yang anda pilih? (H1/H2/H3/H4/H5)")
		var hChoice string
		fmt.Scan(&hChoice)
		hChoice = strings.ToUpper(hChoice)

		targets := model.FindTargets(solver)
		switch hChoice {
		case "H1":
			heuristic = model.H1Manhattan(targets)
		case "H2":
			heuristic = model.H2Euclidean(targets)
		case "H3":
			heuristic = model.H3Chebyshev(targets)
		case "H4":
			heuristic = model.H4MissingTargets()
		case "H5":
			heuristic = model.H5StraightToGoal(targets)
		default:
			fmt.Println("Heuristic tidak valid, menggunakan H1.")
			heuristic = model.H1Manhattan(targets)
		}
	}

	start := time.Now()
	var result model.AlgorithmResult

	switch algo {
	case "UCS":
		result = solver.UCS(state)
	case "GBFS":
		result = solver.GBFS(state, heuristic)
	case "A*":
		result = solver.AStar(state, heuristic)
	case "BFS":
		result = solver.BFS(state)
	default:
		fmt.Println("Algoritma tidak valid.")
		return
	}

	duration := time.Since(start)

	if !result.Solved {
		fmt.Println("Tidak ada solusi yang ditemukan.")
		fmt.Printf(">> Banyak iterasi yang dilakukan: %d iterasi\n", result.Iterations)
		fmt.Printf(">> Waktu eksekusi: %d ms\n", duration.Milliseconds())
		return
	}

	states, moves := solver.GetRoute(result.FinalState)

	fmt.Printf("Solusi Yang Ditemukan : %s\n", string(moves))
	fmt.Printf("Cost dari Solusi : %d\n", result.FinalState.Cost)
	fmt.Printf(">> Waktu eksekusi: %d ms\n", duration.Milliseconds())
	fmt.Printf(">> Banyak iterasi yang dilakukan: %d iterasi\n", result.Iterations)

	fmt.Println(">> Apakah Anda ingin melakukan playback? (Ya/Tidak) :")
	var playback string
	fmt.Scan(&playback)
	if strings.ToLower(playback) == "ya" {
		runPlayback(solver, states, moves)
	}

	fmt.Println(">> Apakah Anda ingin menyimpan solusi? (Ya/Tidak) :")
	var save string
	fmt.Scan(&save)
	if strings.ToLower(save) == "ya" {
		saveSolution(solver, states, moves, result, duration)
	}
}

func runPlayback(solver model.Solver, states []*model.State, moves []rune) {
	fmt.Println("Memulai Playback. Gunakan Panah Kanan (Maju), Panah Kiri (Mundur), ESC (Lompat), Q (Selesai Playback).")
	if err := keyboard.Open(); err != nil {
		fmt.Println("Gagal membuka keyboard listener:", err)
		return
	}
	defer keyboard.Close()

	currentStep := 0
	printStep(solver, states, moves, currentStep)

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			break
		}

		if char == 'q' || char == 'Q' {
			break
		}

		if key == keyboard.KeyArrowRight {
			if currentStep < len(states)-1 {
				currentStep++
				printStep(solver, states, moves, currentStep)
			}
		} else if key == keyboard.KeyArrowLeft {
			if currentStep > 0 {
				currentStep--
				printStep(solver, states, moves, currentStep)
			}
		} else if key == keyboard.KeyEsc {
			keyboard.Close()
			fmt.Print(">> Pada step berapa anda ingin melakukan playback : \n")
			var targetStep int
			fmt.Scan(&targetStep)
			if targetStep >= 0 && targetStep < len(states) {
				currentStep = targetStep
			} else {
				fmt.Println("Step tidak valid.")
			}
			keyboard.Open()
			printStep(solver, states, moves, currentStep)
		}
	}
	fmt.Println("\nPlayback Selesai.")
}

func printStep(solver model.Solver, states []*model.State, moves []rune, step int) {
	fmt.Print("\033[H\033[2J") 
	if step == 0 {
		fmt.Println("Initial")
	} else {
		fmt.Printf("Step %d : %c\n", step, moves[step-1])
	}
	fmt.Println(solver.VisualizeState(*states[step]))
	fmt.Printf("(Step %d dari %d)\n", step, len(states)-1)
}

func saveSolution(solver model.Solver, states []*model.State, moves []rune, result model.AlgorithmResult, duration time.Duration) {
	fmt.Println(">> Nama file untuk menyimpan solusi (contoh: solusi.txt):")
	var saveFilename string
	fmt.Scan(&saveFilename)

	outDir := "../test/solution/"
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		os.MkdirAll(outDir, 0755)
	}

	fullPath := outDir + saveFilename
	file, err := os.Create(fullPath)
	if err != nil {
		fmt.Println("Gagal menyimpan solusi:", err)
		return
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("Solusi Yang Ditemukan : %s\n", string(moves)))
	file.WriteString(fmt.Sprintf("Cost dari Solusi : %d\n", result.FinalState.Cost))
	file.WriteString(fmt.Sprintf("Waktu eksekusi: %d ms\n", duration.Milliseconds()))
	file.WriteString(fmt.Sprintf("Banyak iterasi: %d iterasi\n\n", result.Iterations))

	file.WriteString("Initial\n")
	file.WriteString(solver.VisualizeState(*states[0]))
	file.WriteString("\n")

	for i := 1; i < len(states); i++ {
		file.WriteString(fmt.Sprintf("Step %d : %c\n", i, moves[i-1]))
		file.WriteString(solver.VisualizeState(*states[i]))
		file.WriteString("\n")
	}

	fmt.Printf(">> Solusi disimpan pada %s\n", fullPath)
}
