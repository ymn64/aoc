package main

import (
	_ "embed"
	"fmt"
)

//go:embed input
var stream string

func main() {
	// Part 1
	fmt.Println(countProcessed(4))

	// Part 2
	fmt.Println(countProcessed(14))
}

func countProcessed(markerLen int) int {
	for i := 0; i < len(stream); i++ {
		if isMarker(stream[i : i+markerLen]) {
			return i + markerLen
		}
	}
	return -1
}

func isMarker(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}
