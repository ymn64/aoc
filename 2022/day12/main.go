package main

import (
	"fmt"

	"github.com/devkvlt/aoc/utils"
)

type Map [][]byte

type Pos struct{ x, y int }

func (m Map) isValid(p Pos) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(m) && p.y < len(m[0])
}

type Area map[Pos]struct{}

func (a Area) add(p Pos) {
	a[p] = struct{}{}
}

func (m Map) options(p Pos) Area {
	opts := make(Area)
	dirs := []Pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for _, dir := range dirs {
		o := Pos{p.x + dir.x, p.y + dir.y}
		if m.isValid(o) && int(m[o.x][o.y])-int(m[p.x][p.y]) <= 1 {
			opts.add(o)
		}
	}

	return opts
}

func parse(lines []string) (Map, Pos, Pos, []Pos) {
	m := make(Map, len(lines))
	var start Pos
	var end Pos
	lowest := []Pos{}

	for i := 0; i < len(lines); i++ {
		m[i] = make([]byte, len(lines[0]))

		for j := 0; j < len(lines[0]); j++ {
			m[i][j] = lines[i][j]

			switch lines[i][j] {
			case 'S':
				start = Pos{i, j}
				m[i][j] = 'a'
				lowest = append(lowest, Pos{i, j})
			case 'E':
				end = Pos{i, j}
				m[i][j] = 'z'
			case 'a':
				lowest = append(lowest, Pos{i, j})
			}

		}
	}

	return m, start, end, lowest
}

// https://en.wikipedia.org/wiki/Breadth-first_search
func trek(m Map, currentPos, targetPos Pos) int {
	if currentPos == targetPos {
		return 0
	}

	queue := []Pos{currentPos}
	visited := make(Area)
	visited.add(currentPos)
	steps := 0

	for len(queue) > 0 {
		steps++
		size := len(queue)

		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			opts := m.options(curr)

			for next := range opts {
				if _, ok := visited[next]; !ok {
					if next == targetPos {
						return steps
					}
					queue = append(queue, next)
					visited.add(next)
				}
			}

		}
	}

	return -1
}

func part1(lines []string) {
	m, start, end, _ := parse(lines)

	fmt.Println(trek(m, start, end))
}

func part2(lines []string) {
	m, _, end, lowest := parse(lines)

	best := trek(m, lowest[0], end)

	for i := 1; i < len(lowest); i++ {
		s := trek(m, lowest[i], end)
		if s < best && s != -1 {
			best = s
		}
	}

	fmt.Println(best)
}

func main() {
	lines := utils.ReadLines("input")

	part1(lines) // 423
	part2(lines) // 416
}
