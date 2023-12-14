package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/devkvlt/aoc/utils"
)

var grid []string

func reset() {
	grid = utils.ReadLines("input")
}

func transpose() {
	tr := make([]string, len(grid[0]))
	for i := 0; i < len(grid[0]); i++ {
		s := make([]byte, len(grid))
		for j := 0; j < len(grid); j++ {
			s[j] = grid[j][i]
		}
		tr[i] = string(s)
	}
	grid = tr
}

func tiltLeft() {
	for i, line := range grid {
		parts := strings.Split(line, "#")
		for j := 0; j < len(parts); j++ {
			b := []byte(parts[j])
			slices.SortFunc(b, func(a, b byte) int { return int(b) - int(a) })
			parts[j] = string(b)
		}
		grid[i] = strings.Join(parts, "#")
	}
}

func tiltRight() {
	for i, line := range grid {
		parts := strings.Split(line, "#")
		for j := 0; j < len(parts); j++ {
			b := []byte(parts[j])
			slices.Sort(b)
			parts[j] = string(b)
		}
		grid[i] = strings.Join(parts, "#")
	}
}

func part1() {
	reset()
	transpose()
	tiltLeft()

	sum := 0

	for _, line := range grid {
		for i, ch := range line {
			if ch == 'O' {
				sum += len(line) - i
			}
		}
	}

	fmt.Println(sum)
}

func cycle() {
	transpose()
	tiltLeft()
	transpose()
	tiltLeft()
	transpose()
	tiltRight()
	transpose()
	tiltRight()
}

func key(lines []string) string {
	return strings.Join(lines, "")
}

func part2() {
	reset()

	seen := map[string]bool{key(grid): true}
	grids := [][]string{grid}
	last := 0

	for {
		last++
		cycle()
		if seen[key(grid)] {
			break
		}
		seen[key(grid)] = true
		grids = append(grids, grid)
	}

	first := slices.IndexFunc(grids, func(s []string) bool {
		return slices.Equal(s, grid)
	})

	i := (1000000000-first)%(last-first) + first
	grid = grids[i]
	transpose()

	sum := 0

	for _, line := range grid {
		for i, ch := range line {
			if ch == 'O' {
				sum += len(line) - i
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	part1() // 113525
	part2() // 101292
}
