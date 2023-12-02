package day01

import (
	"strconv"

	"github.com/jarodlam/advent-of-code-2023/pkg/utils"
)

type Solution struct{}

var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func reverseString(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func firstNumber(line string, digits bool, reversed bool) int {
	if reversed {
		line = reverseString(line)
	}

	for pos, char := range line {
		// Look for digits
		n, err := strconv.Atoi(string(char))
		if err == nil {
			return n
		}

		if !digits {
			continue
		}

		// Look for strings
		for numString, n := range numbers {
			var checkString string
			if endPos := pos + len(numString); endPos <= len(line) {
				checkString = line[pos:endPos]
			} else {
				continue
			}

			if reversed {
				checkString = reverseString(checkString)
			}

			if checkString == numString {
				return n
			}
		}
	}
	panic("Can't find first number for line: " + line)
}

func (Solution) Part1(input string) any {
	sum := 0
	utils.ForEachLine(input, func(line string) {
		first := firstNumber(line, false, false)
		last := firstNumber(line, false, true)
		value := first*10 + last
		sum += value
	})

	return sum
}

func (Solution) Part2(input string) any {
	sum := 0
	utils.ForEachLine(input, func(line string) {
		first := firstNumber(line, true, false)
		last := firstNumber(line, true, true)
		value := first*10 + last
		sum += value
	})

	return sum
}
