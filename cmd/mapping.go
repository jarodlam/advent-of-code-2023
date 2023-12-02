package main

import (
	"strconv"

	"github.com/jarodlam/advent-of-code-2023/pkg/day01"
	"github.com/jarodlam/advent-of-code-2023/pkg/day02"
	"github.com/jarodlam/advent-of-code-2023/pkg/utils"
)

func newSolution(day int) utils.Solution {
	switch day {
	case 1:
		return day01.Solution{}
	case 2:
		return day02.Solution{}
	}
	panic("Invalid day: " + strconv.Itoa(day))
}
