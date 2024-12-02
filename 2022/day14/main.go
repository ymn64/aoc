package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/devkvlt/aoc"
)

var lines = aoc.ReadLines("input")

type Map [][]byte

func newMap(width, height int) Map {
	m := make(Map, height)

	for i := 0; i < height; i++ {
		m[i] = make([]byte, width)
		for j := 0; j < width; j++ {
			m[i][j] = '.'
		}
	}

	return m
}

// func debugMap(m Map, xStart, xEnd, yStart, yEnd int) {
// 	for i := yStart; i < min(len(m), yEnd+1); i++ {
// 		for j := xStart; j < min(len(m[0]), xEnd+1); j++ {
// 			fmt.Printf("%c", m[i][j])
// 		}
// 		fmt.Println()
// 	}
// }

type Pos struct{ x, y int }

func parseCoordsList(lines []string) [][]Pos {
	coordsList := make([][]Pos, len(lines))

	for i := 0; i < len(lines); i++ {
		coords := strings.Split(lines[i], " -> ")
		coordsList[i] = make([]Pos, len(coords))

		for j := 0; j < len(coords); j++ {
			xy := strings.Split(coords[j], ",")
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])
			coordsList[i][j] = Pos{x, y}
		}
	}

	return coordsList
}

func drawRocks(m Map, coordsList [][]Pos) {
	for _, coords := range coordsList {
		for i := 0; i < len(coords)-1; i++ {
			xStart, xEnd := coords[i].x, coords[i+1].x
			yStart, yEnd := coords[i].y, coords[i+1].y

			if xStart > xEnd {
				xStart, xEnd = xEnd, xStart
			}

			if yStart > yEnd {
				yStart, yEnd = yEnd, yStart
			}

			for x := xStart; x <= xEnd; x++ {
				for y := yStart; y <= yEnd; y++ {
					m[y][x] = '#'
				}
			}
		}
	}
}

// drawSand returns true if sand was drawn and false otherwise.
func drawSand(m Map) bool {
	x, y := 500, 0

	mustContinue := func() bool {
		return 0 <= x && x < len(m[0]) && y < len(m) && m[y][x] == '.'
	}

	for y < len(m) {
		y++
		if mustContinue() {
			continue
		}

		x--
		if mustContinue() {
			continue
		}

		x += 2
		if mustContinue() {
			continue
		}

		x--
		y--
		if y != len(m)-1 && m[y][x] != '#' {
			m[y][x] = 'o'
			return true
		}

		break
	}

	return false
}

func drawFloor(m Map) {
	y := 0

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if m[i][j] == '#' && i > y {
				y = i
			}
		}
	}

	for j := 0; j < len(m[0]); j++ {
		m[y+2][j] = '#'
	}
}

func part1() {
	m := newMap(1000, 1000)
	coordsList := parseCoordsList(lines)
	drawRocks(m, coordsList)

	units := 0

	ok := drawSand(m)
	for ok {
		units++
		ok = drawSand(m)
	}

	fmt.Println(units)
}

func part2() {
	m := newMap(1000, 1000)
	coordsList := parseCoordsList(lines)
	drawRocks(m, coordsList)
	drawFloor(m)

	units := 0

	for m[0][500] != 'o' {

		units++
		drawSand(m)
	}

	fmt.Println(units)
}

func main() {
	part1() // 757
	part2() // 24943
}
