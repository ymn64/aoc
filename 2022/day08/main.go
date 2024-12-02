package main

import (
	"fmt"
	"strconv"

	"github.com/devkvlt/aoc"
)

var forest [][]int

func main() {
	// Parse trees
	lines := aoc.ReadLines("input")
	forest = make([][]int, len(lines))
	for i := 0; i < len(lines); i++ {
		for _, char := range lines[i] {
			x, _ := strconv.Atoi(string(char))
			forest[i] = append(forest[i], x) // TODO: use len(lines[i]) to avoid append
		}
	}

	// Part 1
	totalVisible := 0
	for i, line := range forest {
		for j := range line {
			if isVisibleFromWest(i, j) ||
				isVisibleFromEast(i, j) ||
				isVisibleFromNorth(i, j) ||
				isVisibleFromSouth(i, j) {
				totalVisible++
			}
		}
	}
	fmt.Println(totalVisible)

	// Part 2
	bestScore := 0
	for i, row := range forest {
		for j := range row {
			s := score(i, j)
			if s > bestScore {
				bestScore = s
			}
		}
	}
	fmt.Println(bestScore)
}

// func isRowVizible(i, j int) bool {
// 	if j == 0 || j == len(forest[0])-1 {
// 		return true
// 	}
//
// 	tree := forest[i][j]
// 	rows := len(forest)
// 	cols := len(forest[0])
//
// 	vis := true
// 	for k := 0; k < j; k++ {
// 		if forest[i][k] >= tree {
// 			vis = false
// 			break
// 		}
// 	}
// 	if vis {
// 		return true
// 	}
//
// 	vis = true
// 	for k := j + 1; k < cols; k++ {
// 		if forest[i][k] >= tree {
// 			vis = false
// 			break
// 		}
// 	}
// 	if vis {
// 		return true
// 	}
// }

func isVisibleFromWest(i, j int) bool {
	if j == 0 {
		return true
	}
	tree := forest[i][j]
	for k := 0; k < j; k++ {
		if forest[i][k] >= tree {
			return false
		}
	}
	return true
}

func isVisibleFromEast(i, j int) bool {
	if j == len(forest[0])-1 {
		return true
	}
	tree := forest[i][j]
	width := len(forest[0])
	for k := j + 1; k < width; k++ {
		if forest[i][k] >= tree {
			return false
		}
	}
	return true
}

func isVisibleFromNorth(i, j int) bool {
	if i == 0 {
		return true
	}
	tree := forest[i][j]
	for k := 0; k < i; k++ {
		if forest[k][j] >= tree {
			return false
		}
	}
	return true
}

func isVisibleFromSouth(i, j int) bool {
	if i == len(forest)-1 {
		return true
	}
	tree := forest[i][j]
	height := len(forest)
	for k := i + 1; k < height; k++ {
		if forest[k][j] >= tree {
			return false
		}
	}
	return true
}

func score(i, j int) int {
	rows := len(forest)
	cols := len(forest[0])
	if i == 0 || i == rows-1 || j == 0 || j == cols-1 {
		return 0
	}
	var scoreLeft, scoreRight, scoreUp, scoreDown int
	tree := forest[i][j]
	for k := j - 1; k >= 0; k-- {
		scoreLeft++
		if forest[i][k] >= tree {
			break
		}
	}
	for k := j + 1; k < cols; k++ {
		scoreRight++
		if forest[i][k] >= tree {
			break
		}
	}
	for k := i - 1; k >= 0; k-- {
		scoreUp++
		if forest[k][j] >= tree {
			break
		}
	}
	for k := i + 1; k < rows; k++ {
		scoreDown++
		if forest[k][j] >= tree {
			break
		}
	}
	return scoreLeft * scoreRight * scoreUp * scoreDown
}
