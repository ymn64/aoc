package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/devkvlt/aoc"
)

type Equation struct {
	val  int
	nums []int
}

var equations []Equation

func init() {
	lines := aoc.Lines("input")
	n := len(lines)
	equations = make([]Equation, n)
	for i := 0; i < n; i++ {
		parts := strings.Split(lines[i], ": ")
		nums := strings.Fields(parts[1])
		equations[i].val = aoc.Atoi(parts[0])
		equations[i].nums = make([]int, len(nums))
		for j, num := range nums {
			equations[i].nums[j] = aoc.Atoi(num)
		}
	}
}

func part1() {
	var possible func(int, []int) bool
	possible = func(val int, nums []int) bool {
		n := len(nums)
		x := nums[n-1]
		if n == 1 {
			return val == x
		}
		var a, b bool
		if val > x {
			a = possible(val-x, nums[:n-1])
		}
		if val%x == 0 {
			b = possible(val/x, nums[:n-1])
		}
		return a || b
	}
	sum := 0
	for _, eq := range equations {
		if possible(eq.val, eq.nums) {
			sum += eq.val
		}
	}
	fmt.Println(sum)
}

func part2() {
	var possible func(int, []int) bool
	possible = func(val int, nums []int) bool {
		n := len(nums)
		x := nums[n-1]
		if n == 1 {
			return val == x
		}
		var a, b, c bool
		if val > x {
			a = possible(val-x, nums[:n-1])
		}
		if val%x == 0 {
			b = possible(val/x, nums[:n-1])
		}
		x_ := strconv.Itoa(x)
		val_ := strconv.Itoa(val)
		m := len(val_) - len(x_)
		if m >= 0 && val_[m:] == x_ {
			sub := val_[:m]
			c = sub != "" && possible(aoc.Atoi(sub), nums[:n-1])
		}
		return a || b || c
	}
	sum := 0
	for _, eq := range equations {
		if possible(eq.val, eq.nums) {
			sum += eq.val
		}
	}
	fmt.Println(sum)
}

func main() {
	part1() // 3312271365652
	part2() // 509463489296712
}
