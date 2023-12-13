package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func off(grid1, grid2 []string) int {
	off := 0
A:
	for i := 0; i < len(grid1); i++ {
		for j := 0; j < len(grid1[i]); j++ {
			if grid1[i][j] != grid2[i][j] {
				off++
				if off == 2 {
					break A
				}
			}
		}
	}
	return off
}

func pointOfIncidence(grid []string, smudge int) int {
	for row := 1; row < len(grid); row++ {

		above := slices.Clone(grid[:row])
		slices.Reverse(above)

		below := grid[row:]

		k := min(len(above), len(below))
		above = above[:k]

		if off(above, below) == smudge {
			return row
		}
	}

	return 0
}

func transpose(grid []string) []string {
	tr := make([]string, len(grid[0]))
	for i := 0; i < len(grid[0]); i++ {
		s := make([]byte, len(grid))
		for j := 0; j < len(grid); j++ {
			s[j] = grid[j][i]
		}
		tr[i] = string(s)
	}
	return tr
}

func solveWith(smudge int) {
	raw, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	chunks := strings.Split(strings.TrimSpace(string(raw)), "\n\n")

	sum := 0

	for _, chunk := range chunks {
		grid := strings.Split(chunk, "\n")
		sum += 100*pointOfIncidence(grid, smudge) + pointOfIncidence(transpose(grid), smudge)
	}

	fmt.Println(sum)
}

func main() {
	solveWith(0) // 35210
	solveWith(1) // 31974
}
