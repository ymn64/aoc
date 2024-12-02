package main

import (
	"container/heap"
	"fmt"

	"github.com/devkvlt/aoc"
)

// https://pkg.go.dev/container/heap#example-package-PriorityQueue

type State struct {
	heatloss int
	pos      Pos
	dir      Pos
	straight int
}

type PQ []*State

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].heatloss < pq[j].heatloss
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x any) {
	*pq = append(*pq, x.(*State))
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	state := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return state
}

var grid = aoc.ReadLines("input")

var (
	height = len(grid)
	width  = len(grid[0])
)

type Pos struct{ x, y int }

var (
	none  = Pos{0, 0}
	up    = Pos{0, -1}
	down  = Pos{0, 1}
	left  = Pos{-1, 0}
	right = Pos{1, 0}
)

func isValid(p Pos) bool {
	return p.x >= 0 && p.x < width && p.y >= 0 && p.y < height
}

func advance(p, dir Pos) Pos {
	return Pos{p.x + dir.x, p.y + dir.y}
}

func split(dir Pos) []Pos {
	if dir == none {
		return []Pos{right, down}
	}
	return []Pos{{dir.y, dir.x}, {-dir.y, -dir.x}}
}

type Key struct {
	pos      Pos
	dir      Pos
	straight int
}

func solveWith(minStraight, maxStraight int) {
	start := Pos{0, 0}
	end := Pos{width - 1, height - 1}

	seen := map[Key]bool{}

	pq := PQ{}
	heap.Push(&pq, &State{heatloss: 0, pos: start, dir: none, straight: 0})

	for len(pq) > 0 {
		s := heap.Pop(&pq).(*State)

		if s.pos == end && s.straight >= minStraight {
			fmt.Println(s.heatloss)
			break
		}

		key := Key{s.pos, s.dir, s.straight}

		if seen[key] {
			continue
		}

		seen[key] = true

		if s.straight < maxStraight && s.dir != none {
			next := advance(s.pos, s.dir)
			if isValid(next) {
				heap.Push(&pq, &State{s.heatloss + int(grid[next.y][next.x]) - 48, next, s.dir, s.straight + 1})
			}
		}

		if s.straight >= minStraight || s.dir == none {
			for _, dir := range split(s.dir) {
				next := advance(s.pos, dir)
				if isValid(next) {
					heap.Push(&pq, &State{s.heatloss + int(grid[next.y][next.x]) - 48, next, dir, 1})
				}
			}
		}

	}
}

func main() {
	solveWith(1, 3)  // 758
	solveWith(4, 10) // 892
}
