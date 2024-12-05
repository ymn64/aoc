package main

import (
	"fmt"

	"github.com/devkvlt/aoc"
)

var grid = aoc.Lines("input")

var (
	height = len(grid)
	width  = len(grid[0])
)

type Pos struct{ x, y int }

var (
	up    = Pos{0, -1}
	down  = Pos{0, 1}
	left  = Pos{-1, 0}
	right = Pos{1, 0}
)

func isValid(p Pos) bool         { return 0 <= p.x && p.x < width && 0 <= p.y && p.y < height }
func advance(p Pos, dir Pos) Pos { return Pos{p.x + dir.x, p.y + dir.y} }
func rotateCW(dir Pos) Pos       { return Pos{dir.y, dir.x} }
func rotateACW(dir Pos) Pos      { return Pos{-dir.y, -dir.x} }

type Ray struct{ pos, dir Pos }

func countEnergized(start Ray) int {
	energized := map[Pos]bool{}
	seen := map[Ray]bool{}
	stack := []Ray{start}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		pos := curr.pos
		dir := curr.dir

		if isValid(pos) && !seen[curr] {
			seen[curr] = true
			energized[pos] = true

			switch grid[pos.y][pos.x] {
			case '.':
				stack = append(stack, Ray{advance(pos, dir), dir})

			case '|':
				if dir == up || dir == down {
					stack = append(stack, Ray{advance(pos, dir), dir})
				} else {
					stack = append(stack, Ray{advance(pos, up), up}, Ray{advance(pos, down), down})
				}

			case '-':
				if dir == left || dir == right {
					stack = append(stack, Ray{advance(pos, dir), dir})
				} else {
					stack = append(stack, Ray{advance(pos, left), left}, Ray{advance(pos, right), right})
				}

			case '/':
				newDir := rotateACW(dir)
				stack = append(stack, Ray{advance(pos, newDir), newDir})

			case '\\':
				newDir := rotateCW(dir)
				stack = append(stack, Ray{advance(pos, newDir), newDir})
			}
		}

	}

	return len(energized)
}

func part1() {
	start := Ray{Pos{0, 0}, right}
	fmt.Println(countEnergized(start))
}

func part2() {
	best := 0

	for x := 0; x < width; x++ {
		start1 := Ray{Pos{x, 0}, down}
		start2 := Ray{Pos{x, height - 1}, up}

		best = max(best, countEnergized(start1), countEnergized(start2))
	}

	for y := 0; y < height; y++ {
		start1 := Ray{Pos{0, y}, right}
		start2 := Ray{Pos{width - 1, y}, left}

		best = max(best, countEnergized(start1), countEnergized(start2))
	}

	fmt.Println(best)
}

func main() {
	part1() // 6361
	part2() // 6701
}
