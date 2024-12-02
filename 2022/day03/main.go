package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/devkvlt/aoc"
)

func main() {
	lines := aoc.ReadLines("input")

	// Part 1
	var total1 int
	for _, line := range lines {
		var common rune
		first := line[:len(line)/2]
		second := line[len(line)/2:]
		for _, item := range first {
			if strings.ContainsRune(second, item) {
				common = item
			}
		}
		total1 += priority(common)

	}
	fmt.Println(total1)

	// Part 2
	var total2 int
	for i := 0; i < len(lines); i += 3 {
		var common rune
		first := lines[i]
		second := lines[i+1]
		third := lines[i+2]
		for _, item := range first {
			if strings.ContainsRune(second, item) && strings.ContainsRune(third, item) {
				common = item
			}
		}
		total2 += priority(common)
	}
	fmt.Println(total2)
}

func priority(item rune) int {
	if unicode.IsLower(item) {
		return int(item - 'a' + 1)
	}
	return int(item - 'A' + 27)
}
