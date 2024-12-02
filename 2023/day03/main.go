package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/devkvlt/aoc"
)

type char interface {
	rune | byte
}

func isDigit[T char](c T) bool {
	return '0' <= c && c <= '9'
}

func isSymbol[T char](c T) bool {
	return c != '.' && !isDigit(c)
}

func capture(line string, j int) int {
	str := string(line[j])

	for k := j + 1; k < len(line); k++ {
		c := line[k]
		if !isDigit(c) {
			break
		}
		str = str + string(c)
	}
	for k := j - 1; k >= 0; k-- {
		c := line[k]
		if !isDigit(c) {
			break
		}
		str = string(c) + str
	}

	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}

	return num
}

func main() {
	lines := aoc.ReadLines("input")

	// Part 1
	sum := 0

	for i, line := range lines {
		for j, c := range line {
			if isSymbol(c) {
				for k := i - 1; k <= i+1; k++ {
					for l := j - 1; l <= j+1; l++ {
						if isDigit(lines[k][l]) && (l == j-1 || !isDigit(lines[k][l-1])) {
							sum += capture(lines[k], l)
						}
					}
				}
			}
		}
	}

	fmt.Println(sum)

	// Part 2
	sum2 := 0

	for i, line := range lines {
		for j, c := range line {
			if c == '*' {
				ratio := 1
				count := 0

				for k := i - 1; k <= i+1; k++ {
					for l := j - 1; l <= j+1; l++ {
						if isDigit(lines[k][l]) && (l == j-1 || !isDigit(lines[k][l-1])) && count < 2 {
							ratio *= capture(lines[k], l)
							count++
						}
					}
				}

				if count == 2 {
					sum2 += ratio
				}
			}
		}
	}

	fmt.Println(sum2)
}
