package main

import (
	"fmt"

	"github.com/devkvlt/aoc"
)

var (
	grid   = aoc.Lines("input")
	height = len(grid)
	width  = len(grid[0])
)

type Pos struct{ x, y int }

func valid(p Pos) bool {
	return 0 <= p.x && p.x < width && 0 <= p.y && p.y < height
}

func part1() {
	antinodes := map[Pos]bool{}
	for y1 := 0; y1 < height; y1++ {
		for x1 := 0; x1 < width; x1++ {
			if grid[y1][x1] == '.' {
				continue
			}
			for y2 := y1; y2 < height; y2++ {
				start := 0
				if y2 == y1 {
					start = x1 + 1 // Avoid comparing the same position
				}
				for x2 := start; x2 < width; x2++ {
					if grid[y1][x1] == grid[y2][x2] {
						dx := x1 - x2
						dy := y1 - y2
						a1 := Pos{x1 + dx, y1 + dy}
						if valid(a1) {
							antinodes[a1] = true
						}
						a2 := Pos{x2 - dx, y2 - dy}
						if valid(a2) {
							antinodes[a2] = true
						}
					}
				}
			}
		}
	}
	fmt.Println(len(antinodes))
}

func part2() {
	antinodes := map[Pos]bool{}
	for y1 := 0; y1 < height; y1++ {
		for x1 := 0; x1 < width; x1++ {
			if grid[y1][x1] == '.' {
				continue
			}
			for y2 := y1; y2 < height; y2++ {
				start := 0
				if y2 == y1 {
					start = x1 + 1 // Avoid comparing the same position
				}
				for x2 := start; x2 < width; x2++ {
					if grid[y1][x1] == grid[y2][x2] {
						dx := x1 - x2
						dy := y1 - y2
						a1 := Pos{x1, y1}
						for valid(a1) {
							antinodes[a1] = true
							a1.x += dx
							a1.y += dy
						}
						a2 := Pos{x2, y2}
						for valid(a2) {
							antinodes[a2] = true
							a2.x -= dx
							a2.y -= dy
						}
					}
				}
			}
		}
	}
	fmt.Println(len(antinodes))
}

func main() {
	part1() // 259
	part2() // 927
}
