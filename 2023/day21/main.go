package main

import (
	"fmt"

	"github.com/devkvlt/aoc"
)

var grid = aoc.Lines("input")

var size = len(grid)

type Pos struct{ x, y int }

var dirs = []Pos{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func isValid(p Pos) bool     { return p.x >= 0 && p.x < size && p.y >= 0 && p.y < size }
func advance(p, dir Pos) Pos { return Pos{p.x + dir.x, p.y + dir.y} }

type PosDist struct {
	pos  Pos
	dist int
}

func part1() {
	steps := 64
	reached := map[Pos]bool{}
	seen := map[Pos]bool{}
	start := Pos{size / 2, size / 2}
	queue := []PosDist{{start, 0}}

	for len(queue) > 0 {
		p := queue[0].pos
		d := queue[0].dist
		queue = queue[1:]

		if isValid(p) && grid[p.y][p.x] != '#' && !seen[p] {
			seen[p] = true

			if d <= steps && d%2 == 0 {
				reached[p] = true
			}

			for _, dir := range dirs {
				queue = append(queue, PosDist{advance(p, dir), d + 1})
			}
		}

	}

	fmt.Println(len(reached))
}

func main() {
	part1() // 3651
}
