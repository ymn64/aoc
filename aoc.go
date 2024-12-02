package aoc

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadLines(path string) []string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}

func Chunks(path string) [][]string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	chunks := strings.Split(strings.TrimSpace(string(bytes)), "\n\n")
	split := make([][]string, len(chunks))
	for i, chunk := range chunks {
		split[i] = strings.Split(chunk, "\n")
	}
	return split
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func Hextoi(hex string) int {
	i64, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(i64)
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(nums []int) int {
	lcm := 1
	for i := 0; i < len(nums); i++ {
		lcm *= nums[i] / GCD(lcm, nums[i])
	}
	return lcm
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func SameSign(x, y int) bool {
	return (x >= 0 && y >= 0) || (x <= 0 && y <= 0)
}
