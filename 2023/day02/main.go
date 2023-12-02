package main

import (
	"fmt"
	"strings"

	"github.com/devkvlt/aoc/utils"
)

type subset = map[string]int

func parseGame(line string) []subset {
	g := []subset{}

	for _, sub := range strings.Split(strings.Split(line, ": ")[1], "; ") {
		colors := strings.Split(sub, ", ")
		m := map[string]int{}
		g = append(g, m)
		for _, color := range colors {
			c := ""
			n := 0
			fmt.Sscanf(color, "%d %s", &n, &c)
			m[c] = n
		}
	}

	return g
}

func isPossible(g []subset) bool {
	for _, sub := range g {
		if sub["red"] > 12 || sub["green"] > 13 || sub["blue"] > 14 {
			return false
		}
	}

	return true
}

func power(g []subset) int {
	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	for _, sub := range g {
		if sub["red"] > maxRed {
			maxRed = sub["red"]
		}
		if sub["green"] > maxGreen {
			maxGreen = sub["green"]
		}
		if sub["blue"] > maxBlue {
			maxBlue = sub["blue"]
		}

	}

	return maxRed * maxGreen * maxBlue
}

func main() {
	lines := utils.ReadLines("input")

	games := make([][]subset, len(lines))

	for i, line := range lines {
		games[i] = parseGame(line)
	}

	// Part 1
	result := 0
	for i, g := range games {
		if isPossible(g) {
			result += i + 1
		}
	}
	fmt.Println(result)

	// Part 2
	result2 := 0
	for _, g := range games {
		result2 += power(g)
	}
	fmt.Println(result2)
}
