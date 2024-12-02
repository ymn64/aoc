package main

import (
	"fmt"
	"strings"

	"github.com/devkvlt/aoc"
)

type Rule struct {
	category byte
	min      int
	max      int
	next     string
}

func parseWorkflow(line string) (string, []Rule) {
	fields := strings.Split(line, "{")
	name := fields[0]
	conds := strings.Split(fields[1][:len(fields[1])-1], ",")

	workflow := make([]Rule, len(conds))

	for i, cond := range conds {
		if !strings.ContainsRune(cond, ':') {
			workflow[i] = Rule{next: cond}
			continue
		}
		category := cond[0]
		split := strings.Split(cond, ":")
		threshold := aoc.Atoi(split[0][2:])
		next := split[1]
		mi := 1
		ma := 4000
		if cond[1] == '>' {
			mi = threshold + 1
		} else {
			ma = threshold - 1
		}
		workflow[i] = Rule{category, mi, ma, next}
	}

	return name, workflow
}

type Part map[byte]int

func parsePart(line string) Part {
	p := make(Part)
	line = line[1 : len(line)-1]
	rs := strings.Split(line, ",")
	for _, r := range rs {
		k := strings.Split(r, "=")[0][0]
		v := strings.Split(r, "=")[1]
		p[k] = aoc.Atoi(v)
	}
	return p
}

var (
	workflows map[string][]Rule
	parts     []Part
)

func init() {
	chunks := aoc.Chunks("input")

	workflows = make(map[string][]Rule, len(chunks[0]))
	parts = make([]Part, len(chunks[1]))

	for _, line := range chunks[0] {
		id, wf := parseWorkflow(line)
		workflows[id] = wf
	}

	for i, line := range chunks[1] {
		parts[i] = parsePart(line)
	}
}

func part1() {
	total := 0

	for _, part := range parts {
		next := "in"

		for {
			if next == "A" {
				total += part['x'] + part['m'] + part['a'] + part['s']
				break
			} else if next == "R" {
				break
			}

		RulesLoop:
			for _, rule := range workflows[next] {
				if rule.min == 0 {
					next = rule.next
				} else {
					rating := part[rule.category]
					if rule.min <= rating && rating <= rule.max {
						next = rule.next
						break RulesLoop
					}

				}
			}
		}
	}

	fmt.Println(total)
}

func main() {
	part1() // 377025
}
