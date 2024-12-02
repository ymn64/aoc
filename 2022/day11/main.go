package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/devkvlt/aoc"
)

type monkey struct {
	items       []int
	inspections int
	div         int
	op          string
	factor      string
	trueTo      int
	falseTo     int
}

func (m monkey) worry() int {
	x := m.items[0]

	if m.op == "+" {
		if m.factor == "old" {
			return x + x
		} else {
			f, err := strconv.Atoi(m.factor)
			if err != nil {
				log.Fatal(err)
			}
			return x + f
		}
	} else {
		if m.factor == "old" {
			return x * x
		} else {
			f, err := strconv.Atoi(m.factor)
			if err != nil {
				log.Fatal(err)
			}
			return x * f
		}
	}
}

func (m monkey) to(x int) int {
	if x%m.div == 0 {
		return m.trueTo
	}
	return m.falseTo
}

func parseMonkey(lines []string) monkey {
	m := monkey{}

	for _, it := range strings.Split(lines[1][18:], ", ") {
		item, err := strconv.Atoi(it)
		if err != nil {
			log.Fatal(err)
		}
		m.items = append(m.items, item)
	}

	fmt.Sscanf(lines[2][13:], "new = old %s %s", &m.op, &m.factor)

	div, err := strconv.Atoi(lines[3][21:])
	if err != nil {
		log.Fatal(err)
	}
	m.div = div

	trueTo, err := strconv.Atoi(lines[4][29:])
	if err != nil {
		log.Fatal(err)
	}
	m.trueTo = trueTo

	falseTo, err := strconv.Atoi(lines[5][30:])
	if err != nil {
		log.Fatal(err)
	}
	m.falseTo = falseTo

	return m
}

func main() {
	lines := aoc.ReadLines("input")

	n := (len(lines) + 1) / 7

	monkeys := make([]monkey, n)

	commDiv := 1

	for i := 0; i < n; i++ {
		monkeys[i] = parseMonkey(lines[7*i : 7*i+6])
		commDiv *= monkeys[i].div
	}

	// rounds := 20 // Part 1
	rounds := 10000 // Part 2

	for i := 0; i < rounds; i++ {
		for j := 0; j < n; j++ {
			for len(monkeys[j].items) > 0 {
				worryLevel := monkeys[j].worry()
				// worryLevel /= 3 // Part 1
				worryLevel %= commDiv // Part 2
				to := monkeys[j].to(worryLevel)
				monkeys[to].items = append(monkeys[to].items, worryLevel)

				monkeys[j].items = monkeys[j].items[1:]
				monkeys[j].inspections++
			}
		}
	}

	// Monkey business
	a := monkeys[0].inspections
	b := monkeys[1].inspections

	for i := 2; i < len(monkeys); i++ {
		if monkeys[i].inspections > a {
			b = a
			a = monkeys[i].inspections
		} else if monkeys[i].inspections > b {
			b = monkeys[i].inspections
		}
	}

	fmt.Println(a * b) // Part 1: 151312, Part 2: 51382025916
}
