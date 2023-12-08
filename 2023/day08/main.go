package main

import (
	"fmt"
	"strings"

	"github.com/devkvlt/aoc/utils"
)

var lines = utils.ReadLines("input")

func parseNodes() (map[string][2]string, []string) {
	nodes := map[string][2]string{}

	startingNodes := []string{}

	for _, line := range lines[2:] {
		nodes[line[:3]] = [2]string{line[7:10], line[12:15]}
		if strings.HasSuffix(line[:3], "A") {
			startingNodes = append(startingNodes, line[:3])
		}
	}

	return nodes, startingNodes
}

func part1() {
	nodes, _ := parseNodes()

	steps := 0

	current := "AAA"

	for current != "ZZZ" {
		for _, direction := range lines[0] {
			if direction == 'L' {
				current = nodes[current][0]
			} else {
				current = nodes[current][1]
			}
			steps++
		}
	}

	fmt.Println(steps)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(nums []int) int {
	lcm := 1
	for i := 0; i < len(nums); i++ {
		lcm *= nums[i] / gcd(lcm, nums[i])
	}
	return lcm
}

func part2() {
	nodes, startingNodes := parseNodes()

	stepsList := make([]int, len(startingNodes))

	for i := 0; i < len(startingNodes); i++ {
		current := startingNodes[i]

		for !strings.HasSuffix(current, "Z") {
			for _, direction := range lines[0] {
				if direction == 'L' {
					current = nodes[current][0]
				} else {
					current = nodes[current][1]
				}
				stepsList[i]++
			}
		}

	}

	fmt.Println(lcm(stepsList))
}

func main() {
	part1()
	part2()
}
