package day01

import (
	"bufio"
	"strconv"
	"strings"
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
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func firstNumber(line string, digits bool, reversed bool) int {
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
	scanner := bufio.NewScanner(strings.NewReader(input))
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		first := firstNumber(line, false, false)
		last := firstNumber(reverseString(line), false, true)

		value := first*10 + last
		sum += value
	}

	return sum
}

func (Solution) Part2(input string) any {
	scanner := bufio.NewScanner(strings.NewReader(input))
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		first := firstNumber(line, true, false)
		last := firstNumber(reverseString(line), true, true)

		value := first*10 + last
		sum += value
	}

	return sum
}
