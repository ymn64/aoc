package main

import (
	"fmt"
	"slices"

	"github.com/devkvlt/aoc"
)

var (
	grid   = aoc.Lines("input")
	height = len(grid)
	width  = len(grid[0])
)

func valid(x, y int) bool {
	return 0 <= x && x < width && 0 <= y && y < height
}

func count1(x, y int) int {
	if grid[y][x] != 'X' {
		return 0
	}
	c := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if !valid(x+3*dx, y+3*dy) {
				continue
			}
			word := make([]byte, 4)
			for i := range 4 {
				word[i] = grid[y+i*dy][x+i*dx]
			}
			if string(word) == "XMAS" {
				c++
			}
		}
	}
	return c
}

func part1() {
	sum := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			sum += count1(x, y)
		}
	}
	fmt.Println(sum)
}

// M.M    M.S    S.M    S.S
// .A.    .A.    .A.    .A.
// S.S    M.S    S.M    M.M
//
// MMSS   MSSM   SMMS   SSMM

func count2(x, y int) int {
	if grid[y][x] != 'A' {
		return 0
	}
	word := string([]byte{grid[y-1][x-1], grid[y-1][x+1], grid[y+1][x+1], grid[y+1][x-1]})
	if slices.Contains([]string{"MMSS", "MSSM", "SMMS", "SSMM"}, word) {
		return 1
	}
	return 0
}

func part2() {
	sum := 0
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			sum += count2(x, y)
		}
	}
	fmt.Println(sum)
}

func main() {
	part1() // 2397
	part2() // 1824
}
