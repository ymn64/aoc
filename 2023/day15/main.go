package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/devkvlt/aoc"
)

var steps = strings.Split(aoc.ReadLines("input")[0], ",")

func hash(s string) int {
	n := 0
	for _, ch := range s {
		n += int(ch)
		n *= 17
		n %= 256
	}
	return n
}

func part1() {
	sum := 0

	for _, s := range steps {
		sum += hash(s)
	}

	fmt.Println(sum)
}

func part2() {
	boxes := make([][]string, 256)
	focalLengths := map[string]int{}

	for _, s := range steps {
		if s[len(s)-1] == '-' {
			label := s[:len(s)-1]
			bi := hash(label)
			li := slices.Index(boxes[bi], label)

			if li != -1 {
				boxes[bi] = append(boxes[bi][:li], boxes[bi][li+1:]...)
			}

			delete(focalLengths, label)
		} else {
			label := s[:len(s)-2]
			bi := hash(label)
			li := slices.Index(boxes[bi], label)

			if li == -1 {
				boxes[bi] = append(boxes[bi], label)
			}
			focalLengths[label] = aoc.Atoi(s[len(s)-1:])
		}
	}

	sum := 0

	for i, box := range boxes {
		for j, lens := range box {
			sum += (i + 1) * (j + 1) * focalLengths[lens]
		}
	}

	fmt.Println(sum)
}

func main() {
	part1() // 515974
	part2() // 265894
}
