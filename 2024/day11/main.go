package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/devkvlt/aoc"
)

var stones = strings.Fields(aoc.Lines("input")[0])

func trim(stone string) string {
	for len(stone) > 1 && stone[0] == '0' {
		stone = stone[1:]
	}
	return stone
}

var cache = map[string]int{}

func count(stone string, n int) int {
	if n == 0 {
		return 1
	}
	key := fmt.Sprintf("%s:%d", stone, n)
	c, ok := cache[key]
	if ok {
		return c
	}
	if stone == "0" {
		c = count("1", n-1)
	} else if l := len(stone); l%2 == 0 {
		c = count(stone[:l/2], n-1) + count(trim(stone[l/2:]), n-1)
	} else {
		c = count(strconv.Itoa(aoc.Atoi(stone)*2024), n-1)
	}
	cache[key] = c
	return c
}

func blink(n int) {
	sum := 0
	for _, stone := range stones {
		sum += count(stone, n)
	}
	fmt.Println(sum)
}

func main() {
	blink(25) // Part 1: 183435
	blink(75) // Part 2: 218279375708592
}
