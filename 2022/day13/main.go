package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/devkvlt/aoc/utils"
)

// lazyUnpack unpacks the top level of a packet.
func lazyUnpack(s string) []string {
	data := []string{}

	s = s[1 : len(s)-1]

	current := ""
	leftBrackets := 0
	rightBrackets := 0

	for len(s) > 0 {
		switch s[:1] {
		case ",":
			if leftBrackets > 0 {
				current += s[:1]
			} else {
				if current != "" {
					data = append(data, current)
				}
				current = ""
			}

		case "[":
			leftBrackets++
			current += s[:1]

		case "]":
			rightBrackets++
			current += s[:1]
			if leftBrackets == rightBrackets {
				data = append(data, current)
				current = ""
				leftBrackets = 0
				rightBrackets = 0
			}

		default:
			current += s[:1]
		}

		s = s[1:]
	}

	if current != "" {
		data = append(data, current)
	}

	return data
}

// cmp returns -1 if the packets are in the right order, 1 if they're in the
// wrong order and 0 if it's a tie.
func cmp(l, r string) int {
	left := lazyUnpack(l)
	right := lazyUnpack(r)

	for i := 0; i < len(left); i++ {
		if i >= len(right) {
			return 1
		}

		isLeftInt := !strings.HasPrefix(left[i], "[")
		isRightInt := !strings.HasPrefix(right[i], "[")

		if isLeftInt && isRightInt {
			a, _ := strconv.Atoi(left[i])
			b, _ := strconv.Atoi(right[i])

			if a < b {
				return -1
			} else if a > b {
				return 1
			}
		} else {
			if isLeftInt {
				left[i] = "[" + left[i] + "]"
			}

			if isRightInt {
				right[i] = "[" + right[i] + "]"
			}

			result := cmp(left[i], right[i])
			if result != 0 {
				return result
			}
		}
	}

	if len(left) < len(right) {
		return -1
	}

	return 0
}

func part1(lines []string) {
	sum := 0

	for i := 0; i < len(lines); i += 3 {
		result := (cmp(lines[i], lines[i+1]))
		if result == -1 {
			sum += i/3 + 1
		}
	}

	fmt.Println(sum)
}

func part2(lines []string) {
	packets := []string{"[[2]]", "[[6]]"}

	for _, line := range lines {
		if line != "" {
			packets = append(packets, line)
		}
	}

	slices.SortFunc(packets, cmp)

	a := slices.Index(packets, "[[2]]") + 1
	b := slices.Index(packets, "[[6]]") + 1

	fmt.Println(a * b)
}

func main() {
	lines := utils.ReadLines("input")

	part1(lines) // 6568
	part2(lines) // 19493
}
