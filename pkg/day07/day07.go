package day07

import (
	"github.com/jarodlam/advent-of-code-2023/pkg/utils"
	"sort"
	"strconv"
	"strings"
)

type Solution struct{}

type hand struct {
	cards []rune
	bid   int
}

var cardValues = map[rune]int{
	'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
	'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
}

func (h hand) String() string {
	return string(h.cards) + " " + strconv.Itoa(h.bid)
}

func handType(h hand) int {
	// Sort into buckets
	buckets := make(map[rune]int)
	for _, r := range h.cards {
		if _, ok := buckets[r]; !ok {
			buckets[r] = 0
		}
		buckets[r] += 1
	}

	// Count combos
	combos := map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0}
	for _, r := range buckets {
		combos[r] += 1
	}

	// Look for combos
	if combos[5] == 1 {
		return 6 // Five of a kind
	} else if combos[4] == 1 && combos[1] == 1 {
		return 5 // Four of a kind
	} else if combos[3] == 1 && combos[2] == 1 {
		return 4 // Full house
	} else if combos[3] == 1 && combos[1] == 2 {
		return 3 // Three of a kind
	} else if combos[2] == 2 && combos[1] == 1 {
		return 2 // Two pair
	} else if combos[2] == 1 && combos[1] == 3 {
		return 1 // One pair
	} else {
		return 0 // High card
	}
}

func handLessThan(h1, h2 hand) bool {
	// Rank by hand type
	ht1 := handType(h1)
	ht2 := handType(h2)
	if ht1 != ht2 {
		return ht1 < ht2
	}

	// Rank by card
	for i := 0; i < 5; i++ {
		val1 := cardValues[h1.cards[i]]
		val2 := cardValues[h2.cards[i]]
		if val1 != val2 {
			return val1 < val2
		}
	}

	return false
}

func (Solution) Part1(input string) any {
	// Parse
	hands := make([]hand, 0)
	utils.ForEachLine(input, func(line string) {
		cardsStr, bidStr, _ := strings.Cut(line, " ")
		cards := []rune(cardsStr)
		bid, _ := strconv.Atoi(bidStr)
		hands = append(hands, hand{cards, bid})
	})

	// Sort by type
	sort.Slice(hands, func(i int, j int) bool {
		return handLessThan(hands[i], hands[j])
	})

	// Calculate total winnings
	total := 0
	for i, h := range hands {
		total += h.bid * (i + 1)
	}

	return total
}

func (Solution) Part2(input string) any {
	return 0
}
