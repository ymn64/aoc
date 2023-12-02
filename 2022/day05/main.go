package main

import (
	"fmt"
	"os"
	"strings"
)

type step struct{ quantity, from, to int }

type stack []byte

func main() {
	bytes, _ := os.ReadFile("input")
	parts := strings.Split(string(bytes), "\n\n")

	// Parse stacks
	lines := strings.Split(parts[0], "\n")
	stackCount := (len(lines[0]) + 1) / 4 // each stack occupies 4 chars
	stacks := make([]stack, stackCount)

	for i := len(lines) - 2; i >= 0; i-- { // ignore line with stack numbers as the info is redundant
		line := lines[i]
		for j := 1; j < len(line); j += 4 {
			char := line[j]
			if 'A' <= char && char <= 'Z' {
				stacks[j/4] = append(stacks[j/4], char) // ** start indexing stacks from 0 instead of from 1
			}
		}
	}

	// Parse steps
	lines = strings.Split(strings.TrimSpace(parts[1]), "\n")
	steps := make([]step, len(lines))

	for i, line := range lines {
		var quantity, from, to int
		fmt.Sscanf(line, "move %d from %d to %d", &quantity, &from, &to)
		steps[i] = step{quantity, from - 1, to - 1} // -1 because **
	}

	// Part 1 VRWBSFZWM
	stacks1 := make([]stack, stackCount)
	copy(stacks1, stacks)

	for _, s := range steps {
		for i := 0; i < s.quantity; i++ {
			n := len(stacks1[s.from]) - 1
			crate := stacks1[s.from][n]
			stacks1[s.to] = append(stacks1[s.to], crate)
			stacks1[s.from] = stacks1[s.from][:n]
		}
	}

	msg1 := make([]byte, stackCount)
	for i, s := range stacks1 {
		msg1[i] = s[len(s)-1]
	}
	fmt.Println(string(msg1))

	// Part 2 RBTWJWMCF
	stacks2 := make([]stack, stackCount)
	copy(stacks2, stacks)

	for _, s := range steps {
		n := len(stacks2[s.from]) - s.quantity
		crates := stacks2[s.from][n:]
		stacks2[s.to] = append(stacks2[s.to], crates...)
		stacks2[s.from] = stacks2[s.from][:n]
	}

	msg2 := make([]byte, stackCount)
	for i, s := range stacks2 {
		msg2[i] = s[len(s)-1]
	}
	fmt.Println(string(msg2))
}
