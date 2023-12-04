package day03

import (
	"testing"
)

const (
	testInput = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
)

var sol = Solution{}

func TestPart1(t *testing.T) {
	want := 4361
	got := sol.Part1(testInput)
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 467835
	got := sol.Part2(testInput)
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
