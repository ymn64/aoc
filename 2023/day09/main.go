package main

import (
	"fmt"
	"strings"

	"github.com/devkvlt/aoc/utils"
)

var lines = utils.ReadLines("input")

func isZero(s []int) bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}
	return true
}

func diffs(s []int) []int {
	diffs := make([]int, len(s)-1)
	for i := 0; i < len(s)-1; i++ {
		diffs[i] = s[i+1] - s[i]
	}
	return diffs
}

func part1() {
	sum := 0

	for _, line := range lines {
		nums := strings.Fields(line)

		history := []int{}
		for _, num := range nums {
			history = append(history, utils.Atoi(num))
		}

		last := []int{history[len(history)-1]}

		for !isZero(history) {
			history = diffs(history)
			last = append(last, history[len(history)-1])
		}

		for i := 0; i < len(last); i++ {
			sum += last[i]
		}
	}

	fmt.Println(sum)
}

func part2() {
	sum := 0

	for _, line := range lines {
		nums := strings.Fields(line)

		history := []int{}
		for _, num := range nums {
			history = append(history, utils.Atoi(num))
		}

		first := []int{history[0]}

		for !isZero(history) {
			history = diffs(history)
			first = append(first, history[0])
		}

		for i := 0; i < len(first); i++ {
			e := 1 - i&1<<1 // 1 if i is even, 0 if odd
			sum += e * first[i]
		}
	}

	fmt.Println(sum)
}

func main() {
	part1() // 1798691765
	part2() // 1104
}
