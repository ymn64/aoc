package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var calories []int64
	var sum int64

	for scanner.Scan() {
		if scanner.Text() != "" {
			i, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			sum += i
		} else {
			calories = append(calories, sum)
			sum = 0
		}
	}
	calories = append(calories, sum)

	// Part 1
	var max int64
	for _, v := range calories {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)

	// Part 2
	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})
	top3 := calories[0] + calories[1] + calories[2]
	fmt.Println(top3)
}
