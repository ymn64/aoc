package main

import (
	"fmt"

	"github.com/devkvlt/aoc"
)

type Machine struct{ xA, yA, xB, yB, x, y int }

var machines []Machine

func init() {
	chunks := aoc.Chunks("input")
	n := len(chunks)
	machines = make([]Machine, n)
	for i, chunk := range chunks {
		fmt.Sscanf(chunk[0], "Button A: X+%d, Y+%d", &machines[i].xA, &machines[i].yA)
		fmt.Sscanf(chunk[1], "Button B: X+%d, Y+%d", &machines[i].xB, &machines[i].yB)
		fmt.Sscanf(chunk[2], "Prize: X=%d, Y=%d", &machines[i].x, &machines[i].y)
	}
}

// α = number of times we push A
// β = number of times we push B
// System to solve: (1): x = α*xA + β*xB
//                  (2): y = α*yA + β*yB
// (1) => α*xA = x - β*xB
//     =>    α = (x - β*xB)/xA
// (2) =>           y = (x - β*xB)*yA/xA + β*yB
//	   =>        y*xA = (x - β*xB)*yA + β*yB*xA
//	   =>        y*xA = x*yA - β*xB*yA + β*yB*xA
//	   => y*xA - x*yA = β*(yB*xA - xB*yA)
//
// FINALLY: β = (y*xA - x*yA)/(yB*xA - xB*yA)
//          α = (x - β*xB)/xA
//      price = 3*α + β

func price(m Machine) int {
	x := m.x
	y := m.y
	xA := m.xA
	yA := m.yA
	xB := m.xB
	yB := m.yB
	if (y*xA-x*yA)%(yB*xA-xB*yA) != 0 {
		return 0
	}
	β := (y*xA - x*yA) / (yB*xA - xB*yA)
	if (x-β*xB)%xA != 0 {
		return 0
	}
	α := (x - β*xB) / xA
	return 3*α + β
}

func part1() {
	sum := 0
	for _, m := range machines {
		sum += price(m)
	}
	fmt.Println(sum)
}

func part2() {
	sum := 0
	for _, m := range machines {
		m.x += 10000000000000
		m.y += 10000000000000
		sum += price(m)
	}
	fmt.Println(sum)
}

func main() {
	part1() // 40069
	part2() // 71493195288102
}
