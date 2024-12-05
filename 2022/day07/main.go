package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/devkvlt/aoc"
)

type dir struct {
	parent   *dir
	children []*dir
	size     int
}

func totalSize(d *dir) int {
	total := d.size
	for _, child := range d.children {
		total += totalSize(child)
	}
	return total
}

// list returns a slice of pointers to the directory d and all its descendant
// directories.
func list(d *dir) []*dir {
	x := []*dir{d}
	for _, child := range d.children {
		x = append(x, list(child)...)
	}
	return x
}

func main() {
	lines := aoc.Lines("input")
	root := dir{}
	var cwd *dir

	for _, line := range lines {
		switch {
		case line[:4] == "$ cd":
			arg := line[5:]
			switch arg {
			case "/":
				cwd = &root
			case "..":
				cwd = cwd.parent
			default:
				child := &dir{parent: cwd}
				cwd.children = append(cwd.children, child)
				cwd = child
			}
		case line == "$ ls":
			continue
		case line[:3] == "dir":
			continue
		default: // file
			size, _ := strconv.Atoi(strings.Split(line, " ")[0])
			cwd.size += size
		}
	}

	// Part 1
	allDirs := list(&root)
	s := 0
	for _, d := range allDirs {
		x := totalSize(d)
		if x <= 100000 {
			s += x
		}
	}
	fmt.Println(s)

	// Part 2
	unused := 70000000 - totalSize(&root)
	minToDelete := 30000000 - unused

	allSizes := make([]int, len(allDirs))
	for i, d := range allDirs {
		allSizes[i] = totalSize(d)
	}
	sort.Ints(allSizes)
	for _, s := range allSizes {
		if s >= minToDelete {
			fmt.Println(s)
			break
		}
	}
}
