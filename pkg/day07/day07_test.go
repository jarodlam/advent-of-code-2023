package day07

import (
	"testing"
)

const (
	testInput = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
)

var sol = Solution{}

func TestPart1(t *testing.T) {
	want := 6440
	got := sol.Part1(testInput)
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 0
	got := sol.Part2(testInput)
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
