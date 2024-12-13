// NOTE: Part 2 is absolutely disgusting!

package main

import (
	"fmt"

	"github.com/devkvlt/aoc"
)

var (
	grid   [][]byte
	height int
	width  int
)

func init() {
	lines := aoc.Lines("input")
	height = len(lines)
	width = len(lines[0])
	grid = make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}
}

type Pos struct{ x, y int }

func valid(p Pos) bool    { return 0 <= p.x && p.x < width && 0 <= p.y && p.y < height }
func at(p Pos) byte       { return grid[p.y][p.x] }
func up(p Pos) Pos        { return Pos{p.x, p.y - 1} }
func down(p Pos) Pos      { return Pos{p.x, p.y + 1} }
func left(p Pos) Pos      { return Pos{p.x - 1, p.y} }
func right(p Pos) Pos     { return Pos{p.x + 1, p.y} }
func upLeft(p Pos) Pos    { return Pos{p.x - 1, p.y - 1} }
func upRight(p Pos) Pos   { return Pos{p.x + 1, p.y - 1} }
func downRight(p Pos) Pos { return Pos{p.x + 1, p.y + 1} }
func downLeft(p Pos) Pos  { return Pos{p.x - 1, p.y + 1} }

func neighbors(p Pos) []Pos {
	var neighbors []Pos
	for _, p_ := range []Pos{up(p), down(p), left(p), right(p)} {
		if valid(p_) && at(p_) == at(p) {
			neighbors = append(neighbors, p_)
		}
	}
	return neighbors
}

func perimeter(p Pos) int {
	perim := 0
	atp := at(p)
	if p.x == 0 || at(Pos{p.x - 1, p.y}) != atp {
		perim++
	}
	if p.x == width-1 || at(Pos{p.x + 1, p.y}) != atp {
		perim++
	}
	if p.y == 0 || at(Pos{p.x, p.y - 1}) != atp {
		perim++
	}
	if p.y == height-1 || at(Pos{p.x, p.y + 1}) != atp {
		perim++
	}
	return perim
}

func part1() {
	seen := map[Pos]bool{}
	queued := map[Pos]bool{}
	sum := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := Pos{x, y}
			if seen[p] {
				continue
			}
			area := 0
			perim := 0
			queue := []Pos{p}
			queued[p] = true
			for len(queue) > 0 {
				curr := queue[0]
				queue = queue[1:]
				seen[curr] = true
				area++
				perim += perimeter(curr)
				for _, n := range neighbors(curr) {
					if !seen[n] && !queued[n] {
						queued[n] = true
						queue = append(queue, n)
					}
				}
			}
			sum += area * perim
		}
	}
	fmt.Println(sum)
}

func cornerMap(p Pos) map[Pos]int {
	m := map[Pos]int{}
	other := func(p_ Pos) bool { return !valid(p_) || at(p_) != at(p) }

	type X struct {
		dir1      func(Pos) Pos
		dir2      func(Pos) Pos
		cornerPos Pos
	}

	xs := []X{
		{up, left, p},
		{up, right, Pos{p.x + 1, p.y}},
		{down, left, Pos{p.x, p.y + 1}},
		{down, right, Pos{p.x + 1, p.y + 1}},
	}

	for _, x := range xs {
		if other(x.dir1(p)) {
			m[x.cornerPos]++
		}
		if other(x.dir2(p)) {
			m[x.cornerPos]++
		}
		if other(x.dir1(x.dir2(p))) {
			m[x.cornerPos]--
			if m[x.cornerPos] == -1 {
				m[x.cornerPos] = 1
			}
		}
	}

	return m
}

func part2() {
	type Plot struct {
		typ       byte
		area      int
		cornerMap map[Pos]int
	}

	garden := []Plot{}

	seen := map[Pos]bool{}
	queued := map[Pos]bool{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := Pos{x, y}
			if seen[p] {
				continue
			}
			area := 0
			m := map[Pos]int{}
			queue := []Pos{p}
			queued[p] = true
			for len(queue) > 0 {
				curr := queue[0]
				queue = queue[1:]
				seen[curr] = true
				area++
				for pos, count := range cornerMap(curr) {
					if _, ok := m[pos]; !ok {
						m[pos] += count
					}
				}
				for _, n := range neighbors(curr) {
					if !seen[n] && !queued[n] {
						queued[n] = true
						queue = append(queue, n)
					}
				}
			}
			garden = append(garden, Plot{at(p), area, m})
		}
	}

	for i := 0; i < len(garden); i++ {
		for j := i + 1; j < len(garden); j++ {
			if garden[i].typ == garden[j].typ { // different plot but same plot type
				// decrement shared corners
				for p1 := range garden[i].cornerMap {
					for p2 := range garden[j].cornerMap {
						if p1 == p2 {
							garden[i].cornerMap[p1]--
							garden[j].cornerMap[p1]--
						}
					}
				}
			}
		}
	}

	sum := 0
	for _, plot := range garden {
		sides := 0
		for p := range plot.cornerMap {
			sides += plot.cornerMap[p]
		}
		sum += plot.area * sides
	}
	fmt.Println(sum)
}

func main() {
	part1() // 1483212
	part2() // 897062
}
