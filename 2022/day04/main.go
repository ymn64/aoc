package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/devkvlt/aoc"
)

func main() {
	lines := aoc.Lines("input")

	var total1 int
	var total2 int

	for _, line := range lines {
		section1 := strings.Split(line, ",")[0]
		section2 := strings.Split(line, ",")[1]

		min1, _ := strconv.Atoi(strings.Split(section1, "-")[0])
		max1, _ := strconv.Atoi(strings.Split(section1, "-")[1])

		min2, _ := strconv.Atoi(strings.Split(section2, "-")[0])
		max2, _ := strconv.Atoi(strings.Split(section2, "-")[1])

		if (min1 <= min2 && max1 >= max2) || (min2 <= min1 && max2 >= max1) {
			total1++
		}

		if max1 >= min2 && min1 <= max2 {
			total2++
		}
	}

	// Part 1
	fmt.Println(total1)

	// Part 2
	fmt.Println(total2)
}
