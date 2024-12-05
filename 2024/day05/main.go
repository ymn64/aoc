package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/devkvlt/aoc"
)

var (
	rules   map[string][]string
	updates [][]string
)

func init() {
	chunks := aoc.Chunks("input")
	rules = map[string][]string{}
	for _, line := range chunks[0] {
		lr := strings.Split(line, "|")
		rules[lr[0]] = append(rules[lr[0]], lr[1])
	}
	updates = make([][]string, len(chunks[1]))
	for i, line := range chunks[1] {
		updates[i] = strings.Split(line, ",")
	}
}

func cmp(a, b string) int {
	if slices.Contains(rules[a], b) {
		return -1
	}
	return 1
}

func check(u []string) bool {
	for i := 1; i < len(u); i++ {
		if cmp(u[i-1], u[i]) > 0 {
			return false
		}
	}
	return true
}

func part1() {
	sum := 0
	for _, u := range updates {
		if check(u) {
			sum += aoc.Atoi(u[len(u)/2])
		}
	}
	fmt.Println(sum)
}

func part2() {
	sum := 0
	for _, u := range updates {
		if !check(u) {
			slices.SortFunc(u, cmp)
			sum += aoc.Atoi(u[len(u)/2])
		}
	}
	fmt.Println(sum)
}

func main() {
	part1() // 4609
	part2() // 5723
}
