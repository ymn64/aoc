package main

import (
	"fmt"
	"strings"

	"github.com/devkvlt/aoc"
)

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func btoi(c byte) int {
	return int(c - '0')
}

func main() {
	lines := aoc.Lines("input")

	// Part 1
	sum := 0

	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			if isDigit(line[i]) {
				sum += btoi(line[i]) * 10
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if isDigit(line[i]) {
				sum += btoi(line[i])
				break
			}
		}
	}

	fmt.Println(sum) // 56042

	// Part2
	sum2 := 0

	digits := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, line := range lines {
	AdvanceRight:
		for i := 0; i < len(line); i++ {
			for k, v := range digits {
				if strings.HasPrefix(line[i:], k) {
					sum2 += v * 10
					break AdvanceRight
				}
			}
		}
	DvanceLeft:
		for i := len(line) - 1; i >= 0; i-- {
			for k, v := range digits {
				if strings.HasPrefix(line[i:], k) {
					sum2 += v
					break DvanceLeft
				}
			}
		}
	}

	fmt.Println(sum2) // 55358
}
