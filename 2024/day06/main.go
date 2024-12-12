package main

import (
	"fmt"

	"github.com/devkvlt/aoc"
)

var (
	grid    [][]byte
	height  int
	width   int
	start   Pos
	visited []Pos
)

type Pos struct{ x, y int }

func init() {
	lines := aoc.Lines("input")
	height = len(lines)
	width = len(lines[0])
	grid = make([][]byte, height)
	for y := 0; y < height; y++ {
		grid[y] = make([]byte, width)
		for x := 0; x < width; x++ {
			grid[y][x] = lines[y][x]
			ch := lines[y][x]
			if ch == '^' {
				start = Pos{x, y}
			}
		}
	}
}

func valid(p Pos) bool {
	return 0 <= p.x && p.x < width && 0 <= p.y && p.y < height
}

func next(p Pos, dir byte) Pos {
	return map[byte]Pos{
		'^': {p.x, p.y - 1},
		'>': {p.x + 1, p.y},
		'v': {p.x, p.y + 1},
		'<': {p.x - 1, p.y},
	}[dir]
}

func turn(dir byte) byte {
	return map[byte]byte{'^': '>', '>': 'v', 'v': '<', '<': '^'}[dir]
}

func part1() {
	curr := start
	dir := grid[curr.y][curr.x]
	seen := map[Pos]bool{curr: true}
	for valid(curr) {
		next := next(curr, dir)
		if !valid(next) {
			break
		}
		if grid[next.y][next.x] != '#' {
			if !seen[next] {
				visited = append(visited, next)
			}
			seen[next] = true
			curr = next
		} else {
			dir = turn(dir)
		}
	}
	fmt.Println(len(visited) + 1) // +1 to include initial pos
}

type PosDir struct {
	pos Pos
	dir byte
}

func loop(p Pos) int {
	grid[p.y][p.x] = '#'
	curr := start
	dir := grid[curr.y][curr.x]
	seen := map[PosDir]bool{{curr, dir}: true}

	for valid(curr) {
		next := next(curr, dir)
		if !valid(next) {
			break
		}
		if grid[next.y][next.x] == '#' {
			dir = turn(dir)
			continue
		}
		if seen[PosDir{next, dir}] {
			grid[p.y][p.x] = '.'
			return 1
		}
		seen[PosDir{next, dir}] = true
		curr = next
	}
	grid[p.y][p.x] = '.'
	return 0
}

func part2() {
	sum := 0
	for _, p := range visited {
		sum += loop(p)
	}
	fmt.Println(sum)
}

// NOTE: part2 needs part1 to run first because it relies on visited which is
// calculated in part1.
func main() {
	part1() // 4890
	part2() // 1995
}
