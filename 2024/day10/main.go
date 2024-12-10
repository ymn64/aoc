package main

import (
	"fmt"

	"github.com/devkvlt/aoc"
)

type Pos struct{ x, y int }

var (
	grid       = aoc.Lines("input")
	height     = len(grid)
	width      = len(grid[0])
	trailheads []Pos
)

func init() {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if grid[y][x] == '0' {
				trailheads = append(trailheads, Pos{x, y})
			}
		}
	}
}

func elevation(p Pos) byte {
	return grid[p.y][p.x] - '0'
}

func neighbors(p Pos) []Pos {
	var neighbors []Pos
	for _, n := range []Pos{{p.x, p.y + 1}, {p.x, p.y - 1}, {p.x + 1, p.y}, {p.x - 1, p.y}} {
		if n.x >= 0 && n.x < width && n.y >= 0 && n.y < height && elevation(n) == elevation(p)+1 {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

func part1() {
	score := 0
	for _, head := range trailheads {
		queue := []Pos{head}
		seen := map[Pos]bool{}
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			for _, next := range neighbors(curr) {
				if elevation(next) == 9 {
					if !seen[next] {
						seen[next] = true
						score++
					}
				} else {
					queue = append(queue, next)
				}
			}
		}
	}
	fmt.Println(score)
}

func part2() {
	rating := 0
	for _, head := range trailheads {
		queue := []Pos{head}
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			for _, next := range neighbors(curr) {
				if elevation(next) == 9 {
					rating++
				} else {
					queue = append(queue, next)
				}
			}
		}
	}
	fmt.Println(rating)
}

func main() {
	part1() // 517
	part2() // 1116
}
