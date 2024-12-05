package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/devkvlt/aoc"
)

var lines = aoc.Lines("input")

type (
	Hand     = string
	HandType = int
	Card     = byte
)

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func handType(hand Hand) HandType {
	freq := make(map[rune]int)
	maxFreq := 0

	for _, card := range hand {
		freq[card]++
		if freq[card] > maxFreq {
			maxFreq = freq[card]
		}
	}

	switch len(freq) {
	case 5:
		return HighCard
	case 4:
		return OnePair
	case 3:
		if maxFreq == 2 {
			return TwoPair
		}
		return ThreeOfAKind
	case 2:
		if maxFreq == 3 {
			return FullHouse
		}
		return FourOfAKind
	default:
		return FiveOfAKind
	}
}

func cmp1(hand1, hand2 Hand) int {
	diff := handType(hand1) - handType(hand2)
	if diff != 0 {
		return int(diff)
	}

	worth := map[Card]int{
		'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10, '9': 9, '8': 8,
		'7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2,
	}

	for i := 0; i < 5; i++ {
		diff := worth[hand1[i]] - worth[hand2[i]]
		if diff != 0 {
			return diff
		}
	}

	return 0
}

func upgrade(hand Hand) Hand {
	freq := make(map[rune]int)
	maxFreq := 0
	mostFreq := ""

	for _, card := range hand {
		if card != 'J' {
			freq[card]++
			if freq[card] > maxFreq {
				maxFreq = freq[card]
				mostFreq = string(card)
			}
		}
	}

	return strings.ReplaceAll(hand, "J", mostFreq)
}

func cmp2(hand1, hand2 Hand) int {
	diff := handType(upgrade(hand1)) - handType(upgrade(hand2))
	if diff != 0 {
		return int(diff)
	}

	worth := map[Card]int{
		'A': 14, 'K': 13, 'Q': 12, 'T': 10, '9': 9, '8': 8,
		'7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2,
		'J': 1,
	}

	for i := 0; i < 5; i++ {
		diff := worth[hand1[i]] - worth[hand2[i]]
		if diff != 0 {
			return diff
		}
	}

	return 0
}

type Pair struct {
	hand Hand
	bid  int
}

func solveWith(cmpFunc func(a, b Hand) int) {
	list := make([]Pair, len(lines))

	for i, line := range lines {
		list[i] = Pair{
			strings.Fields(line)[0],
			aoc.Atoi(strings.Fields(line)[1]),
		}
	}

	slices.SortFunc(list, func(a, b Pair) int {
		return cmpFunc(a.hand, b.hand)
	})

	winnings := 0

	for i, pair := range list {
		winnings += (i + 1) * pair.bid
	}

	fmt.Println(winnings)
}

func main() {
	solveWith(cmp1) // 251121738
	solveWith(cmp2) // 251421071
}
