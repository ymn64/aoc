package main

import (
	"fmt"

	"github.com/devkvlt/aoc/utils"
)

func main() {
	guide := utils.ReadLines("input")

	// Part 1
	var total1 int
	for _, hint := range guide {
		theirs := int(hint[0] - 'A' + 1)
		mine := int(hint[2] - 'X' + 1)
		total1 += score(theirs, mine)
	}
	fmt.Println(total1)

	// Part 2
	var total2 int
	for _, hint := range guide {
		theirs := int(hint[0] - 'A' + 1)
		outcome := int(hint[2]) - int('Y')
		mine := (theirs + outcome) % 3
		if mine == 0 {
			mine = 3
		}
		total2 += score(theirs, mine)
	}
	fmt.Println(total2)
}

func score(theirs, mine int) int {
	x := (theirs - mine + 3) % 3 // x==2 => win, x==1 => loss, x==0 => draw
	f := func(x int) int {
		return (9*x*x - 15*x + 6) / 2 // f(2)=6, f(1)=0, f(0)=3, god bless Lagrange
	}
	return mine + f(x)
}
