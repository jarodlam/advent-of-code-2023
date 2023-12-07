package day06

import (
	"testing"
)

const (
	testInput = `Time:      7  15   30
Distance:  9  40  200`
)

var sol = Solution{}

func TestPart1(t *testing.T) {
	want := 288
	got := sol.Part1(testInput)
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 71503
	got := sol.Part2(testInput)
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
