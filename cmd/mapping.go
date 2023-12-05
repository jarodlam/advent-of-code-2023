package main

import (
	"strconv"

	"github.com/jarodlam/advent-of-code-2023/pkg/day01"
	"github.com/jarodlam/advent-of-code-2023/pkg/day02"
	"github.com/jarodlam/advent-of-code-2023/pkg/day03"
	"github.com/jarodlam/advent-of-code-2023/pkg/day04"
	"github.com/jarodlam/advent-of-code-2023/pkg/day05"
	"github.com/jarodlam/advent-of-code-2023/pkg/utils"
)

func newSolution(day int) utils.Solution {
	switch day {
	case 1:
		return day01.Solution{}
	case 2:
		return day02.Solution{}
	case 3:
		return day03.Solution{}
	case 4:
		return day04.Solution{}
	case 5:
		return day05.Solution{}
	}
	panic("Invalid day: " + strconv.Itoa(day))
}
