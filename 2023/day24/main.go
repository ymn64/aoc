package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/devkvlt/aoc/utils"
)

type Hailstone struct{ x, y, z, vx, vy, vz float64 }

var hailstones []Hailstone

func init() {
	lines := utils.ReadLines("input")

	hailstones = make([]Hailstone, len(lines))

	for i, line := range lines {
		fmt.Sscanf(
			line,
			"%f, %f, %f @ %f, %f, %f",
			&hailstones[i].x,
			&hailstones[i].y,
			&hailstones[i].z,
			&hailstones[i].vx,
			&hailstones[i].vy,
			&hailstones[i].vz,
		)
	}
}

func willCollide(h1, h2 Hailstone) bool {
	if h1.vy*h2.vx == h1.vx*h2.vy { // This basically checks for a == α (parallel lines)
		return false
	}

	a := h1.vy / h1.vx
	b := h1.y - a*h1.x

	α := h2.vy / h2.vx
	β := h2.y - α*h2.x

	x := (b - β) / (α - a)
	y := a*x + b

	if (x < h1.x && h1.vx > 0) || (x > h1.x && h1.vx < 0) || (x < h2.x && h2.vx > 0) || (x > h2.x && h2.vx < 0) {
		return false
	}

	mi := 200000000000000.0
	ma := 400000000000000.0
	if x < mi || x > ma || y < mi || y > ma {
		return false
	}

	return true
}

func part1() {
	collisions := 0

	for i, h1 := range hailstones {
		for _, h2 := range hailstones[i+1:] {
			if willCollide(h1, h2) {
				collisions++
			}
		}
	}

	fmt.Println(collisions)
}

func part2() {
	script := "import sympy\n"
	script += "x, y, z, vx, vy, vz = sympy.symbols('x, y, z, vx, vy, vz')\n"
	script += "system = [\n"
	for _, h := range hailstones[:4] {
		script += fmt.Sprintf("  (x - %.0f)*(%.0f - vy) - (y - %.0f)*(%.0f - vx),\n", h.x, h.vy, h.y, h.vx)
		script += fmt.Sprintf("  (y - %.0f)*(%.0f - vz) - (z - %.0f)*(%.0f - vy),\n", h.y, h.vz, h.z, h.vy)
	}
	script += "]\n"
	script += "solutions = sympy.solve(system)\n"
	script += "assert len(solutions) == 1\n"
	script += "s = solutions[0]\n"
	script += "print(s[x] + s[y] + s[z])"

	cmd := exec.Command("python3", "-c", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {
	part1() // 16665
	part2() // 769840447420960
}
