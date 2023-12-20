package main

import (
	"fmt"
	"strings"

	"github.com/devkvlt/aoc/utils"
)

type Pulse int8

const (
	Low  Pulse = 0
	High Pulse = 1
)

// // DEBUG:
// func (p Pulse) String() string {
// 	if p == Low {
// 		return "Low"
// 	}
// 	return "High"
// }

type Signal struct {
	pulse Pulse
	src   string
	dst   string
}

type Module interface {
	ID() string
	Dsts() []string
	Process(src string, input Pulse)
}

type FlipFlop struct {
	id   string
	dsts []string
	on   *bool
}

func (f FlipFlop) ID() string     { return f.id }
func (f FlipFlop) Dsts() []string { return f.dsts }

func NewFlipFlop(id string, dsts []string) FlipFlop {
	on := false
	return FlipFlop{id: id, dsts: dsts, on: &on}
}

func (f FlipFlop) Process(src string, input Pulse) {
	if input == Low {
		*f.on = !(*f.on)
		output := Low
		if *f.on {
			output = High
		}
		for _, dst := range f.dsts {
			queue = append(queue, Signal{output, f.ID(), dst})
			// fmt.Println(f.ID(), ": Send", output, "to", dst) // DEBUG:
		}
	}
}

type Conjunction struct {
	id     string
	dsts   []string
	memory *map[string]Pulse
}

func NewConjunction(id string, dsts []string) Conjunction {
	return Conjunction{id: id, dsts: dsts, memory: &map[string]Pulse{}}
}

func (c Conjunction) ID() string     { return c.id }
func (c Conjunction) Dsts() []string { return c.dsts }

func (c Conjunction) Process(src string, input Pulse) {
	(*c.memory)[src] = input
	output := Low
	for _, pulse := range *c.memory {
		if pulse == Low {
			output = High
			break
		}
	}
	for _, dst := range c.dsts {
		queue = append(queue, Signal{output, c.ID(), dst})
		// fmt.Println(c.ID(), ": Send", output, "to", dst) // DEBUG:
	}
}

type Broadcaster struct {
	id   string
	dsts []string
}

func NewBroadcaster(id string, dsts []string) Broadcaster {
	return Broadcaster{id: id, dsts: dsts}
}

func (b Broadcaster) ID() string     { return b.id }
func (b Broadcaster) Dsts() []string { return b.dsts }

func (c Broadcaster) Process(src string, input Pulse) {
	for _, dst := range c.dsts {
		queue = append(queue, Signal{input, c.ID(), dst})
		// fmt.Println(c.ID(), ": Send", input, "to", dst) // DEBUG:
	}
}

var queue []Signal

var modules = map[string]Module{}

func reset() {
	lines := utils.ReadLines("input")

	for _, line := range lines {
		var mod Module

		fields := strings.Split(line, " -> ")

		id := fields[0]
		dsts := strings.Split(fields[1], ", ")

		switch id[0] {
		case '%':
			id = id[1:]
			mod = NewFlipFlop(id, dsts)
		case '&':
			id = id[1:]
			mod = NewConjunction(id, dsts)
		default:
			mod = NewBroadcaster(id, dsts)
		}

		modules[id] = mod
	}

	// Populate the memory for Conjunctions.
	for id, mod := range modules {
		for _, dst := range mod.Dsts() {
			if c, ok := modules[dst].(Conjunction); ok {
				(*c.memory)[id] = Low
			}
		}
	}

	// // DEBUG:
	// for name, mod := range modules {
	// 	if m, ok := mod.(Conjunction); ok {
	// 		fmt.Println(name, mod.Dsts(), (*m.memory))
	// 	} else if m, ok := mod.(FlipFlop); ok {
	// 		fmt.Println(name, mod.Dsts(), (*m.on))
	// 	} else {
	// 		fmt.Println(name, mod.Dsts())
	// 	}
	// }
}

func part1() {
	reset()

	lows := 0
	highs := 0

	for i := 0; i < 1000; i++ {
		queue = []Signal{{Low, "button", "broadcaster"}} // Push da button

		for len(queue) > 0 {
			signal := queue[0]
			queue = queue[1:]

			pulse := signal.pulse
			src := signal.src
			dst := signal.dst

			if pulse == High {
				highs++
			} else {
				lows++
			}

			// // DEBUG:
			// if dst == "output" {
			// 	continue
			// }

			mod, ok := modules[dst]
			if ok {
				mod.Process(src, pulse)
			}
		}
	}

	fmt.Println(lows * highs)
}

// rx gets its input from a single module which happens to be a conjunction.
// For this module to produce a Low pulse, all the modules that feed into it
// must send a High pulse at the same time.
// If we assume that these modules periodically output a High pulse, then we can
// calculate their period and then calculate the LCM of these periods to find
// the number of presses it will take for all of them to produce a High pulse at
// the same time, and that will be the answer to part 2.
func part2() {
	reset()

	// The one conjunction that outputs to rx.
	beforeRX := ""
	for id, mod := range modules {
		for _, dst := range mod.Dsts() {
			if dst == "rx" {
				beforeRX = id
				break
			}
		}
	}

	// The number of modules that output to beforeRX. This is used to make a
	// breaking condition in the loop below.
	n := 0
	for _, mod := range modules {
		for _, dst := range mod.Dsts() {
			if dst == beforeRX {
				n++
			}
		}
	}

	// List of minimum presses required for each of the modules that output to
	// beforeRX to produce a High pulse.
	pressesList := []int{}

	presses := 0

	for len(pressesList) < n {
		presses++
		queue = []Signal{{Low, "button", "broadcaster"}} // Push da button

		for len(queue) > 0 {
			signal := queue[0]
			queue = queue[1:]

			pulse := signal.pulse
			src := signal.src
			dst := signal.dst

			if dst == beforeRX && pulse == High {
				pressesList = append(pressesList, presses)
			}

			mod, ok := modules[dst]
			if ok {
				mod.Process(src, pulse)
			}
		}
	}

	fmt.Println(utils.LCM(pressesList))
}

func main() {
	part1() // 896998430
	part2() // 236095992539963
}
