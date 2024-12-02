package main

import (
	"fmt"
	"strings"

	"github.com/devkvlt/aoc"
)

type Set = map[string]struct{}

func setFromSlice(s []string) Set {
	set := make(Set)
	for _, v := range s {
		set[v] = struct{}{}
	}
	return set
}

func intersect(a, b Set) Set {
	set := make(Set)
	for e := range a {
		if _, ok := b[e]; ok {
			set[e] = struct{}{}
		}
	}
	return set
}

func got(line string) int {
	parts := strings.Split(strings.Split(line, ":")[1], "|")
	have := setFromSlice(strings.Fields(parts[0]))
	winning := setFromSlice(strings.Fields(parts[1]))

	return len(intersect(have, winning))
}

func pow2(n int) int {
	p := 1
	for i := 0; i < n; i++ {
		p *= 2
	}
	return p
}

func part1(lines []string) {
	points := 0

	for i := 0; i < len(lines); i++ {
		got := got(lines[i])
		if got != 0 {
			points += pow2(got - 1)
		}
	}

	fmt.Println(points)
}

func part2(lines []string) {
	n := len(lines)
	copies := make([]int, n)
	sum := n

	for i := 0; i < n; i++ {
		copies[i]++
		got := got(lines[i])
		for j := 0; j < copies[i]; j++ {
			for k := 1; k <= got && i+k < n; k++ {
				copies[i+k]++
				sum++
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	lines := aoc.ReadLines("input")

	part1(lines)
	part2(lines)
}
