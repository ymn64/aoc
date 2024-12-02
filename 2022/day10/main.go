package main

import (
	"fmt"
	"strconv"

	"github.com/devkvlt/aoc"
)

func part1(lines []string) int {
	strength := 0

	x := 1
	cycle := 1

	add := func() {
		switch cycle {
		case 20, 60, 100, 140, 180, 220:
			strength += cycle * x
		}
	}

	for _, line := range lines {
		cycle++
		add()

		if line != "noop" {
			v, _ := strconv.Atoi(line[5:])
			x += v
			cycle++
			add()
		}
	}

	return strength
}

func part2(lines []string) {
	x := 1
	cycle := 1
	screen := ""

	draw := func() {
		pos := (cycle - 1) % 40

		if x-1 <= pos && pos <= x+1 {
			screen += "#"
		} else {
			screen += "."
		}
	}

	for _, line := range lines {
		draw()
		cycle++

		if line != "noop" {
			draw()
			v, _ := strconv.Atoi(line[5:])
			x += v
			cycle++
		}
	}

	for i := 0; i < 6; i++ {
		fmt.Println(screen[i*40 : (i+1)*40])
	}
}

func main() {
	lines := aoc.ReadLines("input")
	fmt.Println(part1(lines)) // 12980
	part2(lines)              // BRJLFULP
}
