// NOTE: Very challenging problem. Brute-forcing part 1 with sets of positions
// takes 2 or 3 seconds but that's not doable in part 2. Instead the problem
// should be solved on paper with elementary geometry and a little bit of
// trickery. The key trick in my solution is the use of interval merging. Good
// luck!

package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/devkvlt/aoc/utils"
)

type Pos struct{ x, y int }

func abs(x int) int {
	return max(x, -x)
}

func dist(a, b Pos) int {
	return abs(b.x-a.x) + abs(b.y-a.y)
}

type Sensor struct {
	pos    Pos
	radius int
}

func parseSensors(filename string) []Sensor {
	lines := utils.ReadLines(filename)

	sensors := make([]Sensor, len(lines))

	for i := 0; i < len(lines); i++ {
		fields := strings.Fields(lines[i])
		sensor := Pos{utils.Atoi(fields[2][2 : len(fields[2])-1]), utils.Atoi(fields[3][2 : len(fields[3])-1])}
		beacon := Pos{utils.Atoi(fields[8][2 : len(fields[8])-1]), utils.Atoi(fields[9][2:])}
		sensors[i] = Sensor{sensor, dist(sensor, beacon)}
	}

	return sensors
}

type Interval struct {
	start, end int
}

func emptyInterval(s Sensor, y int) (Interval, bool) {
	r := s.radius - abs(s.pos.y-y)
	if r < 0 {
		return Interval{}, false
	}
	return Interval{s.pos.x - r, s.pos.x + r}, true
}

func mergedIntervals(sensors []Sensor, y int) []Interval {
	intervals := []Interval{}

	for _, sensor := range sensors {
		interval, ok := emptyInterval(sensor, y)
		if ok {
			intervals = append(intervals, interval)
		}
	}

	slices.SortFunc(intervals, func(a, b Interval) int {
		return a.start - b.start
	})

	merged := []Interval{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		prev := merged[len(merged)-1]
		curr := intervals[i]

		if curr.start-1 > prev.end {
			merged = append(merged, curr)
		} else {
			merged[len(merged)-1].end = max(curr.end, prev.end)
		}
	}

	return merged
}

func part1() {
	// sensors := parseSensors("sample")
	// y := 10
	sensors := parseSensors("input")
	y := 2000000

	merged := mergedIntervals(sensors, y)

	total := 0
	for _, i := range merged {
		total += i.end - i.start
	}

	fmt.Println(total)
}

func part2() {
	// sensors := parseSensors("sample")
	// maximum := 20
	sensors := parseSensors("input")
	maximum := 4000000

	tuningFreq := -1

	for y := 0; y < maximum; y++ {
		merged := mergedIntervals(sensors, y)

		merged[0].start = 0
		merged[len(merged)-1].end = maximum

		if len(merged) == 2 {
			tuningFreq = 4000000*(merged[0].end+1) + y
			break
		}
	}

	fmt.Println(tuningFreq)
}

func main() {
	part1() // 5525990
	part2() // 11756174628223
}
