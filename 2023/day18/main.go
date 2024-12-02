package main

import (
	"fmt"
	"strings"

	"github.com/devkvlt/aoc"
)

var lines = aoc.ReadLines("input")

func area(x, y []int, boundary int) int {
	// https://en.wikipedia.org/wiki/Shoelace_formula
	n := len(x)

	area := 0
	for i := 1; i < n-1; i++ {
		area += x[i] * (y[i-1] - y[i+1])
	}
	area += x[0] * (y[n-1] - y[1])
	area += x[n-1] * (y[n-2] - y[0])
	area /= 2
	area = max(area, -area)

	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	interior := area - boundary/2 + 1

	return interior + boundary
}

func parse1(fields []string) (byte, int) {
	dir := fields[0][0]
	steps := aoc.Atoi(fields[1])

	return dir, steps
}

func parse2(fields []string) (byte, int) {
	m := map[byte]byte{'0': 'R', '1': 'D', '2': 'L', '3': 'U'}
	dir := m[fields[2][7]]
	steps := aoc.Hextoi(fields[2][2:7])

	return dir, steps
}

func solveWith(parseFn func([]string) (byte, int)) {
	currX := 0
	currY := 0
	x := []int{currX}
	y := []int{currY}
	boundary := 0

	for _, line := range lines {
		fields := strings.Fields(line)
		dir, steps := parseFn(fields)

		switch dir {
		case 'R':
			currX += steps
		case 'D':
			currY += steps
		case 'L':
			currX -= steps
		case 'U':
			currY -= steps
		}

		x = append(x, currX)
		y = append(y, currY)

		boundary += steps
	}

	fmt.Println(area(x, y, boundary))
}

func main() {
	solveWith(parse1) // 41019
	solveWith(parse2) // 96116995735219
}
