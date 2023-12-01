package pkg

import (
	"strconv"

	"github.com/jarodlam/advent-of-code-2023/pkg/day01"
)

func NewSolution(day int) Solution {
	switch day {
	case 1:
		return day01.Solution{}
	}
	panic("Invalid day: " + strconv.Itoa(day))
}
