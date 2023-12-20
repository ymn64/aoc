package utils

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
