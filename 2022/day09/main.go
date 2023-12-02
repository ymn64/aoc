package main

import (
	"fmt"

	"github.com/devkvlt/aoc/utils"
)

type rope struct {
	knots  []knot
	series []motion
}

type knot struct {
	x, y int
}

type motion struct {
	dir   byte
	count int
}

func main() {
	lines := utils.ReadLines("input")

	series := make([]motion, len(lines))
	for i, line := range lines {
		fmt.Sscanf(line, "%c %d", &series[i].dir, &series[i].count)
	}

	// Part 1
	rope2 := rope{knots: make([]knot, 2), series: series}
	fmt.Println(countVisited(rope2))

	// Part 2
	rope10 := rope{knots: make([]knot, 10), series: series}
	fmt.Println(countVisited(rope10))
}

func countVisited(r rope) int {
	visited := map[knot]struct{}{{0, 0}: {}}

	for _, mo := range r.series {
		for i := 0; i < mo.count; i++ {
			switch mo.dir {
			case 'U':
				r.knots[0].y++
			case 'D':
				r.knots[0].y--
			case 'L':
				r.knots[0].x--
			case 'R':
				r.knots[0].x++
			}

			for j := 1; j < len(r.knots); j++ {
				dx := r.knots[j-1].x - r.knots[j].x
				dy := r.knots[j-1].y - r.knots[j].y

				if dy == 0 && abs(dx) > 1 {
					r.knots[j].x += abs(dx) / dx
				} else if dx == 0 && abs(dy) > 1 {
					r.knots[j].y += abs(dy) / dy
				} else if abs(dx) > 1 || abs(dy) > 1 {
					r.knots[j].x += abs(dx) / dx
					r.knots[j].y += abs(dy) / dy
				}

			}

			visited[r.knots[len(r.knots)-1]] = struct{}{}
		}
	}

	return len(visited)
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
