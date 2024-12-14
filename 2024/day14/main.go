package main

import (
	"fmt"
	"slices"

	"github.com/devkvlt/aoc"
)

type Robot struct{ x, y, vx, vy int }

var (
	robots []Robot
	width  = 101
	height = 103
	// width  = 11
	// height = 7
)

func init() {
	lines := aoc.Lines("input")
	robots = make([]Robot, len(lines))
	for i, line := range lines {
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robots[i].x, &robots[i].y, &robots[i].vx, &robots[i].vy)
	}
}

func move(robots []Robot) {
	for i := 0; i < len(robots); i++ {
		robots[i].x += robots[i].vx
		if robots[i].x > width-1 {
			robots[i].x = robots[i].x % width
		} else if robots[i].x < 0 {
			robots[i].x = width + robots[i].x%width
		}

		robots[i].y += robots[i].vy
		if robots[i].y > height-1 {
			robots[i].y = robots[i].y % height
		} else if robots[i].y < 0 {
			robots[i].y = height + robots[i].y%height
		}
	}
}

func part1() {
	robots := slices.Clone(robots)
	for range 100 {
		move(robots)
	}
	var a, b, c, d int
	for _, r := range robots {
		switch {
		case r.x < width/2 && r.y < height/2:
			a++
		case r.x < width/2 && height/2 < r.y:
			b++
		case width/2 < r.x && r.y < height/2:
			c++
		case width/2 < r.x && height/2 < r.y:
			d++
		}
	}
	fmt.Println(a * b * c * d)
}

func part2() {
	robots := slices.Clone(robots)
	// isLine checks if we have a vertical line of robots (top robot being r).
	// The idea being that such a configuration can't just occur randomly.
	// NOTE: You might have to try values bigger than 7. For my input 7 is the
	// lowest number that gets the correct answer, and it takes about 3s.
	isLine := func(r Robot) bool {
		lineLen := 7
		for dy := 1; dy < lineLen; dy++ {
			i := slices.IndexFunc(robots, func(r_ Robot) bool {
				return r_.x == r.x && r_.y == r.y+dy
			})
			if i == -1 {
				return false
			}
		}
		return true
	}
	isTree := func() bool {
		for _, r := range robots {
			if isLine(r) {
				return true
			}
		}
		return false
	}
	t := 0
	for !isTree() {
		t++
		move(robots)
	}
	fmt.Println(t)
}

func main() {
	part1() // 224554908
	part2() // 6644
}
