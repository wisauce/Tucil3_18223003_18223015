package main

import (
	"IceSlidingPuzzle/model"
	"IceSlidingPuzzle/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SolveResponse struct {
	Success     bool           `json:"success"`
	Message     string         `json:"message"`
	Path        string         `json:"path"`
	TimeMs      int64          `json:"time_ms"`
	Iterations  int            `json:"iterations"`
	States      []*model.State `json:"states"`
	BoardWidth  int            `json:"board_width"`
	BoardHeight int            `json:"board_height"`
	Board       []string       `json:"board"`
	Costs       [][]int        `json:"costs"`
}

func getTargets(solver model.Solver) map[int][2]int {
	targets := make(map[int][2]int)
	for i := 0; i < len(solver.Board); i++ {
		for j := 0; j < len(solver.Board[0]); j++ {
			c := solver.Board[i][j]
			if c >= '0' && c <= '9' {
				targets[int(c-'0')] = [2]int{j, i}
			} else if c == 'O' {
				targets[solver.FinalNumber+1] = [2]int{j, i}
			}
		}
	}
	return targets
}

func solveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	algorithm := r.FormValue("algorithm")
	heuristicName := r.FormValue("heuristic")

	defer func() {
		if r := recover(); r != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(SolveResponse{Success: false, Message: fmt.Sprintf("Format file tidak valid: %v", r)})
		}
	}()

	initialState, solver := utils.ParseReader(file)

	var hFunc model.HeuristicFunc
	targets := getTargets(solver)
	if algorithm == "gbfs" || algorithm == "a*" {
		switch heuristicName {
		case "h1":
			hFunc = model.H1Manhattan(targets)
		case "h2":
			hFunc = model.H2Euclidean(targets)
		case "h3":
			hFunc = model.H3Chebyshev(targets)
		case "h4":
			hFunc = model.H4MissingTargets()
		case "h5":
			hFunc = model.H5StraightToGoal(targets)
		default:
			hFunc = model.H1Manhattan(targets)
		}
	}

	start := time.Now()
	var result model.AlgorithmResult

	switch algorithm {
	case "ucs":
		result = solver.UCS(initialState)
	case "gbfs":
		result = solver.GBFS(initialState, hFunc)
	case "a*":
		result = solver.AStar(initialState, hFunc)
	case "bfs":
		result = solver.BFS(initialState)
	default:
		result = solver.UCS(initialState)
	}

	elapsed := time.Since(start).Milliseconds()

	if !result.Solved {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SolveResponse{Success: false, Message: "Solusi tidak ditemukan!"})
		return
	}

	states, pathRunes := solver.GetRoute(result.FinalState)
	path := string(pathRunes)
	iterations := result.Iterations

	boardStr := make([]string, len(solver.Board))
	for i, row := range solver.Board {
		boardStr[i] = string(row)
	}

	resp := SolveResponse{
		Success:     true,
		Path:        path,
		TimeMs:      elapsed,
		Iterations:  iterations,
		States:      states,
		BoardWidth:  len(solver.Board[0]),
		BoardHeight: len(solver.Board),
		Board:       boardStr,
		Costs:       solver.Costs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/solve", solveHandler)

	fmt.Println("Server Web berjalan di http://localhost:8080")
	fmt.Println("Tekan Ctrl+C untuk berhenti.")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}
