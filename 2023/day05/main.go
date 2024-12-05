package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/devkvlt/aoc"
)

type Range struct {
	dstStart int
	srcStart int
	len      int
}

type Map []Range

func parseMaps(lines []string) []Map {
	maps := []Map{}

	for j := 0; j < len(lines); j++ {
		if strings.HasSuffix(lines[j], " map:") {
			m := Map{}

			for i := j + 1; i < len(lines); i++ {
				if lines[i] == "" {
					break
				}

				fields := strings.Fields(lines[i])

				d, _ := strconv.Atoi(fields[0])
				s, _ := strconv.Atoi(fields[1])
				l, _ := strconv.Atoi(fields[2])

				m = append(m, Range{d, s, l})
			}

			maps = append(maps, m)
		}
	}
	return maps
}

func src2dst(m Map, src int) int {
	for _, r := range m {
		d := src - r.srcStart
		if 0 <= d && d < r.len {
			return r.dstStart + d
		}
	}
	return src
}

func dst2src(m Map, dst int) int {
	for _, r := range m {
		d := dst - r.dstStart
		if 0 <= d && d < r.len {
			return r.srcStart + d
		}
	}
	return dst
}

func part1(lines []string) {
	seeds := []int{}
	for _, v := range strings.Split(lines[0][7:], " ") {
		seed, _ := strconv.Atoi(v)
		seeds = append(seeds, seed)
	}

	maps := parseMaps(lines)

	best := seeds[0]

	for i := 1; i < len(seeds); i++ {
		src := seeds[i]
		for _, m := range maps {
			src = src2dst(m, src)
		}
		if src < best {
			best = src
		}
	}

	fmt.Println(best)
}

type SeedRange struct {
	start int
	count int
}

func inRange(ranges []SeedRange, seed int) bool {
	for _, r := range ranges {
		if r.start <= seed && seed < r.start+r.count {
			return true
		}
	}
	return false
}

func part2(lines []string) {
	fields := strings.Split(lines[0][7:], " ")
	seedRanges := []SeedRange{}
	for i := 0; i < len(fields); i += 2 {
		start, _ := strconv.Atoi(fields[i])
		count, _ := strconv.Atoi(fields[i+1])
		seedRanges = append(seedRanges, SeedRange{start, count})
	}

	maps := parseMaps(lines)

	best := 0

	for {
		dst := best
		for i := len(maps) - 1; i >= 0; i-- {
			dst = dst2src(maps[i], dst)
		}
		if inRange(seedRanges, dst) {
			break
		}
		best++
	}

	fmt.Println(best)
}

func main() {
	lines := aoc.Lines("input")

	part1(lines) // 174137457
	part2(lines) // 1493866
}
