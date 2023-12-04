package day04

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/jarodlam/advent-of-code-2023/pkg/utils"
)

type Solution struct{}

type intSet = map[int]struct{}

func findAllNumbers(s string) intSet {
	numbers := make(intSet)
	re := regexp.MustCompile(`\d+`)
	found := re.FindAllString(s, -1)
	for _, x := range found {
		n, _ := strconv.Atoi(x)
		numbers[n] = struct{}{}
	}
	return numbers
}

func (Solution) Part1(input string) any {
	total := 0
	utils.ForEachLine(input, func(line string) {
		// Parse
		_, numbersStr, _ := strings.Cut(line, ": ")
		winStr, mineStr, _ := strings.Cut(numbersStr, "|")
		nWin := findAllNumbers(winStr)
		nMine := findAllNumbers(mineStr)

		// Set intersect
		count := 0
		for n := range nWin {
			if _, ok := nMine[n]; ok {
				count += 1
			}
		}

		// Add score
		if count > 0 {
			total += 1 << (count - 1)
		}
	})
	return total
}

func (Solution) Part2(input string) any {
	total := 0
	currCard := 0
	tally := make(map[int]int)
	utils.ForEachLine(input, func(line string) {
		currCard += 1

		// Total
		if _, ok := tally[currCard]; !ok {
			tally[currCard] = 1
		}
		total += tally[currCard]

		// Parse
		_, numbersStr, _ := strings.Cut(line, ": ")
		winStr, mineStr, _ := strings.Cut(numbersStr, "|")
		nWin := findAllNumbers(winStr)
		nMine := findAllNumbers(mineStr)

		// Set intersect
		count := 0
		for n := range nWin {
			if _, ok := nMine[n]; ok {
				count += 1
			}
		}

		// Calculate score
		if count == 0 {
			return
		}

		// Tally
		for i := currCard + 1; i <= (currCard + count); i++ {
			if _, ok := tally[i]; !ok {
				tally[i] = 1
			}
			tally[i] += tally[currCard]
		}
	})

	return total
}
