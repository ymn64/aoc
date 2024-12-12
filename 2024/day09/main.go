// NOTE: kinda difficult...

package main

import (
	"fmt"
	"slices"

	"github.com/devkvlt/aoc"
)

var (
	disk []int
	ptr  int
)

func init() {
	for k, v := range aoc.Lines("input")[0] {
		var what int
		if k%2 == 0 {
			what = k / 2
		} else {
			what = -1
		}
		for range int(v - '0') {
			disk = append(disk, what)
		}
	}
	ptr = len(disk) - 1
}

// func DEBUG() {
// 	for _, x := range disk {
// 		if x == -1 {
// 			fmt.Printf(".")
// 		} else {
// 			fmt.Printf("%d", x)
// 		}
// 	}
// 	fmt.Println()
// }

func checksum(disk []int) int {
	sum := 0
	for i, id := range disk {
		if id != -1 {
			sum += i * id
		}
	}
	return sum
}

func part1() {
	disk := slices.Clone(disk)
	for l, r := 0, len(disk)-1; l < r; {
		if disk[r] != -1 && disk[l] == -1 {
			disk[l] = disk[r]
			disk[r] = -1
		}
		if disk[l] != -1 {
			l++
		}
		if disk[r] == -1 {
			r--
		}
	}
	fmt.Println(checksum(disk))
}

func nextFile() (int, int) {
	end := ptr
	for end >= 0 && disk[end] == -1 {
		end--
	}
	if end < 0 {
		return -1, -1
	}

	start := end
	for start >= 0 && disk[start] == disk[end] {
		start--
	}
	start++

	ptr = start - 1
	return start, end
}

func nextBlankStart(size int) int {
	for i := 0; i <= ptr-size+1; i++ {
		isBlank := true
		for j := 0; j < size; j++ {
			if disk[i+j] != -1 {
				isBlank = false
				break
			}
		}
		if isBlank {
			return i
		}
	}
	return -1
}

func part2() {
	var start, end, size int
	var blankStart int
	for {
		start, end = nextFile()
		size = end - start + 1
		blankStart = nextBlankStart(size)
		if start != -1 {
			if blankStart != -1 {
				for m := 0; m < size; m++ {
					disk[blankStart+m] = disk[start]
					disk[end-m] = -1
				}
			}
		} else {
			break
		}
	}
	fmt.Println(checksum(disk))
}

func main() {
	part1() // 6446899523367
	part2() // 6478232739671
}
