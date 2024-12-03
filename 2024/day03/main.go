// TODO: don't use regex.

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func prod(mem string) int {
	sum := 0
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := re.FindAllString(mem, -1)
	for _, match := range matches {
		var x, y int
		fmt.Sscanf(match, "mul(%d,%d)", &x, &y)
		sum += x * y
	}
	return sum
}

func main() {
	bytes, _ := os.ReadFile("input")
	mem := string(bytes)

	// Part 1: 173517243
	fmt.Println(prod(mem))

	// Part 2: 100450138
	sum := 0
	for {
		start := strings.Index(mem, "don't()")
		if start == -1 {
			break
		}
		sum += prod(mem[:start])
		end := strings.Index(mem[start:], "do()")
		if end == -1 {
			break
		}
		mem = mem[start+end+4:]
	}
	fmt.Println(sum)
}
