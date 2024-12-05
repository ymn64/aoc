package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/devkvlt/aoc"
)

var lines = aoc.Lines("input")

func waysToWin(time, record string) int {
	t, _ := strconv.Atoi(time)
	r, _ := strconv.Atoi(record)

	ways := 0

	for v := 1; v < t; v++ {
		if v*(t-v) > r {
			ways++
		}
	}

	return ways
}

func part1() {
	times := strings.Fields(lines[0])[1:]
	records := strings.Fields(lines[1])[1:]

	prod := 1

	for i := 0; i < len(times); i++ {
		prod *= waysToWin(times[i], records[i])
	}

	fmt.Println(prod)
}

func part2() {
	time := strings.Join(strings.Fields(lines[0])[1:], "")
	record := strings.Join(strings.Fields(lines[1])[1:], "")

	fmt.Println(waysToWin(time, record))
}

func main() {
	part1() // 114400
	part2() // 21039729
}
