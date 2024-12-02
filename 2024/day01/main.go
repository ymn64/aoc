package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/devkvlt/aoc"
)

var (
	n     int
	left  []int
	right []int
)

func init() {
	lines := aoc.ReadLines("input")
	n = len(lines)
	left = make([]int, n)
	right = make([]int, n)
	for i, line := range lines {
		f := strings.Fields(line)
		left[i] = aoc.Atoi(f[0])
		right[i] = aoc.Atoi(f[1])
	}
}

func part1() {
	slices.Sort(left)
	slices.Sort(right)
	sum := 0
	for i := 0; i < n; i++ {
		sum += aoc.Abs(left[i] - right[i])
	}
	fmt.Println(sum)
}

func freq(x int, s []int) int {
	f := 0
	for _, y := range s {
		if x == y {
			f++
		}
	}
	return f
}

func part2() {
	sum := 0
	for _, x := range left {
		sum += x * freq(x, right)
	}

	fmt.Println(sum)
}

func main() {
	part1() // 2164381
	part2() // 20719933
}
