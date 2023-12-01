package day01

import (
	"testing"
)

const (
	testInput1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`
	testInput2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
)

var sol = Solution{}

func TestPart1(t *testing.T) {
	want := 142
	got := sol.Part1(testInput1)
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 281
	got := sol.Part2(testInput2)
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
