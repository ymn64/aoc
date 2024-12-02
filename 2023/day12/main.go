package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/devkvlt/aoc"
)

var lines = aoc.ReadLines("input")

var cache = make(map[string]int)

func key(cond string, nums []int) string {
	strNums := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		strNums[i] = strconv.Itoa(nums[i])
	}
	return cond + strings.Join(strNums, ",")
}

func count(cond string, nums []int) int {
	key := key(cond, nums)

	result, ok := cache[key]
	if ok {
		return result
	}

	if cond == "" {
		if len(nums) == 0 {
			return 1
		}
		return 0
	}

	if len(nums) == 0 {
		if strings.ContainsRune(cond, '#') {
			return 0
		}
		return 1
	}

	if cond[0] == '.' || cond[0] == '?' {
		result += count(cond[1:], nums)
	}

	if (cond[0] == '#' || cond[0] == '?') &&
		nums[0] <= len(cond) &&
		!strings.ContainsRune(cond[:nums[0]], '.') &&
		(nums[0] == len(cond) || cond[nums[0]] != '#') {

		k := min(nums[0]+1, len(cond))
		result += count(cond[k:], nums[1:])
	}

	cache[key] = result
	return result
}

func solveWith(repeat int) {
	sum := 0

	for _, line := range lines {
		fields := strings.Fields(line)

		cond := fields[0]

		rawNums := strings.Split(fields[1], ",")
		n := len(rawNums)
		nums := make([]int, n)
		for i := 0; i < n; i++ {
			nums[i] = aoc.Atoi(rawNums[i])
		}

		for i := 1; i < repeat; i++ {
			cond += "?" + fields[0]
			nums = append(nums, nums[:n]...)
		}

		sum += count(cond, nums)
	}

	fmt.Println(sum)
}

func main() {
	solveWith(1) // 6935
	solveWith(5) // 3920437278260
}
