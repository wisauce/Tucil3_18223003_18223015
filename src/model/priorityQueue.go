package model

type Item struct {
  State    *State
  Priority int // UCS: Cost, A*: Cost + heuristic
  index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
  return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
  return pq[i].Priority < pq[j].Priority // MIN
}

func (pq PriorityQueue) Swap(i, j int) {
  pq[i], pq[j] = pq[j], pq[i]
  pq[i].index = i
  pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
  item := x.(*Item)
  item.index = len(*pq)
  *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
  old := *pq
  n := len(old)
  item := old[n-1]
  old[n-1] = nil
  *pq = old[:n-1]
  return item
}