package main

import (
	"fmt"

	"github.com/devkvlt/aoc"
)

type Pos struct{ x, y int }

var (
	universe  = aoc.Lines("input")
	positions []Pos
	emptyRows []int
	emptyCols []int
)

func init() {
	for row := 0; row < len(universe); row++ {
		for col := 0; col < len(universe[0]); col++ {
			if universe[row][col] == '#' {
				positions = append(positions, Pos{col, row})
			}
		}
	}

	for row := 0; row < len(universe); row++ {
		if isEmptyRow(row) {
			emptyRows = append(emptyRows, row)
		}
	}

	for col := 0; col < len(universe[0]); col++ {
		if isEmptyCol(col) {
			emptyCols = append(emptyCols, col)
		}
	}
}

func isEmptyCol(col int) bool {
	for row := 0; row < len(universe); row++ {
		if universe[row][col] == '#' {
			return false
		}
	}
	return true
}

func isEmptyRow(row int) bool {
	for col := 0; col < len(universe[row]); col++ {
		if universe[row][col] == '#' {
			return false
		}
	}
	return true
}

func dist(p1, p2 Pos, expansion int) int {
	x1 := min(p1.x, p2.x)
	x2 := max(p1.x, p2.x)
	y1 := min(p1.y, p2.y)
	y2 := max(p1.y, p2.y)

	dx := 0
	dy := 0

	for _, emptyCol := range emptyCols {
		if x1 < emptyCol && emptyCol < x2 {
			dx += expansion - 1
		}
	}

	for _, emptyRow := range emptyRows {
		if y1 < emptyRow && emptyRow < y2 {
			dy += expansion - 1
		}
	}

	return x2 - x1 + dx + y2 - y1 + dy
}

func solveWith(expantion int) {
	sum := 0
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			sum += dist(positions[i], positions[j], expantion)
		}
	}
	fmt.Println(sum)
}

func main() {
	solveWith(2)       // 9686930
	solveWith(1000000) // 630728425490
}
