// ┏━┏━┓ ┏━🐥┏━┓ ┏━┓ ┏━┓ ┏━┓ ┏━┓ ┏━━━━━━━┓
// ┗━┃ ┗━┛ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┏━━━━━┛
// ┏━┗━━━┓ ┗━┛ ┗━┛ ┃ ┃ ┃ ┃ ┃ ┃ ┗━┛ ┗━━━┓ ┓
// ┏━━━━━┛ ┏━━━━━┓ ┃ ┃ ┗━┛ ┗━┛ ┓ ┏━┓ ┏━┛ ━━
// ┗━━━━━━━┛ ┏━━━┛ ┗━┛ • ┃ ┃ ━━┏━┛ ┗━┛ ┛ ┓
// ┃ ┏━┃ ┏━━━┛ ┏━━━━━━━┓ ┏━┓ ━━┗━┓ ┗━┃ ┓ ┃
// ┃ ┏━┏━┛ ┏━┓ ┗━┓ ┏━━━┛ ┏━┓ ┃ ┛ ┗━━━━━━━┓
// ┓ ━━┗━━━┛ ┗━┓ ┃ ┃ ┏━┓ ┃ ┗━┓ ┏━━━┓ ┏━┓ ┃
// ┗━• ┗━┓ ┗━┏━┛ ┃ ┃ ┃ ┃ ┃ ┏━┛ ┗━┓ ┃ ┃ ┗━┛
// ┗━┓ ┛ ┗━┛ ┗━━━┛ ┗━┛ ┗━┛ ┗━━━━━┛ ┗━┛ • ┗━

package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/devkvlt/aoc"
)

var maze = aoc.Lines("input")

var start Pos

func init() {
A:
	for row := 0; row < len(maze); row++ {
		for col := 0; col < len(maze[0]); col++ {
			if maze[row][col] == 'S' {
				start = Pos{col, row}
				maze[row] = strings.Replace(maze[row], "S", valueOfS(start), 1)
				break A
			}
		}
	}
}

func visualize() {
	char := map[byte]string{
		'|': "┃ ",
		'-': "━━",
		'F': "┏━",
		'J': "┛ ",
		'L': "┗━",
		'7': "┓ ",
		'.': "• ",
		'S': "🐥",
	}
	for row := 0; row < len(maze); row++ {
		for col := 0; col < len(maze[0]); col++ {
			fmt.Printf("%s", char[maze[row][col]])
		}
		fmt.Println()
	}
}

type Pos struct{ x, y int }

func isValid(p Pos) bool {
	return 0 <= p.x && p.x < len(maze[0]) && 0 <= p.y && p.y < len(maze)
}

func at(p Pos) byte {
	return maze[p.y][p.x]
}

func to(p Pos, dir Pos) Pos {
	return Pos{p.x + dir.x, p.y + dir.y}
}

var (
	up    = Pos{0, -1}
	down  = Pos{0, 1}
	left  = Pos{-1, 0}
	right = Pos{1, 0}
)

func reverse(dir Pos) Pos {
	return Pos{-dir.x, -dir.y}
}

var validDirs = map[byte][]Pos{
	'|': {up, down},
	'-': {left, right},
	'7': {left, down},
	'F': {right, down},
	'J': {up, left},
	'L': {up, right},
}

var validChars = map[Pos][]byte{
	up:    {'|', '7', 'F'},
	down:  {'|', 'L', 'J'},
	left:  {'-', 'L', 'F'},
	right: {'-', 'J', '7'},
}

func neighbors(p Pos) []Pos {
	neighbors := []Pos{}
	for _, dir := range validDirs[at(p)] {
		to := to(p, dir)
		if isValid(to) && slices.Contains(validChars[dir], at(to)) {
			neighbors = append(neighbors, to)
		}

	}
	return neighbors
}

func valueOfS(start Pos) string {
	s := "|-LJ7F"

	for _, dir := range []Pos{up, down, left, right} {
		to := to(start, dir)
		if !isValid(to) || !slices.Contains(validChars[dir], at(to)) {
			for _, char := range validChars[reverse(dir)] {
				s = strings.Replace(s, string(char), "", 1)
			}
		}
	}

	return s
}

func loop() map[Pos]bool {
	loop := map[Pos]bool{start: true}
	queue := []Pos{start}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, next := range neighbors(curr) {
			if !loop[next] {
				loop[next] = true
				queue = append(queue, next)
			}
		}
	}
	return loop
}

func part1() {
	fmt.Println(len(loop()) / 2)
}

func part2() {
	loop := loop()
	outside := maps.Clone(loop)

	for row := 0; row < len(maze); row++ {
		within := false
		upward := false

		for col := 0; col < len(maze[0]); col++ {
			char := maze[row][col]
			if !loop[Pos{col, row}] {
				char = '.'
			}

			if char == '|' {
				within = !within
			} else if char == 'L' {
				upward = true
			} else if char == 'F' {
				upward = false
			} else if (char == '7' && upward) || (char == 'J' && !upward) {
				within = !within
			}

			if !within {
				outside[Pos{col, row}] = true
			}
		}

	}

	fmt.Println(len(maze)*len(maze[0]) - len(outside))
}

func main() {
	// visualize()
	part1() // 6931
	part2() // 357
}
