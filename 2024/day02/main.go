package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/devkvlt/aoc"
)

var reports [][]int

func init() {
	lines := aoc.ReadLines("input")
	reports = make([][]int, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		report := make([]int, len(fields))
		for j, f := range fields {
			report[j] = aoc.Atoi(f)
		}
		reports[i] = report
	}
}

func safe(report []int) bool {
	d0 := report[1] - report[0]
	for i := 1; i < len(report); i++ {
		d := report[i] - report[i-1]
		if d == 0 || aoc.Abs(d) > 3 || d*d0 < 0 {
			return false
		}
	}
	return true
}

func part1() {
	sum := 0
	for _, r := range reports {
		if safe(r) {
			sum++
		}
	}
	fmt.Println(sum)
}

func part2() {
	sum := 0
	for _, r := range reports {
		if safe(r) {
			sum++
			continue
		}
		for i := 0; i < len(r); i++ {
			if safe(slices.Delete(slices.Clone(r), i, i+1)) {
				sum++
				break
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	part1() // 591
	part2() // 621
}
