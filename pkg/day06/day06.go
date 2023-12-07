package day06

import (
	"math"
	"strings"

	"github.com/jarodlam/advent-of-code-2023/pkg/utils"
)

type Solution struct{}

/*
Equation for final distance:

	d = b * (T - b) = -b² + T*b

where
- `b` is length of button press
- `T` is race time

Solving for intercept with record distance (`D`):

	D = -b² + T*b
	0 = -b² + T*b - D
	b = (-T ± sqrt(T² - 4D)) / -2

One-liner solution:

	$$n=1+\left\lfloor\frac{-T-\sqrt{T^2-4(D+1)}}{-2}\right\rfloor-\left\lceil\frac{-T+\sqrt{T^2-4(D+1)}}{-2}\right\rceil$$
*/
func possibleWins(intT int, intD int) int {
	T := float64(intT)
	D := float64(intD + 1) // Add 1 to ensure it beats the record

	xMin := int(math.Ceil((-T + math.Sqrt(math.Pow(T, 2)-4*D)) / -2))
	xMax := int(math.Floor((-T - math.Sqrt(math.Pow(T, 2)-4*D)) / -2))
	diff := xMax - xMin + 1 // Add 1 to count boundary

	return diff
}

func (Solution) Part1(input string) any {
	// Parse input
	inputT := make([]int, 0)
	inputD := make([]int, 0)
	utils.ForEachLine(input, func(line string) {
		if len(inputT) == 0 {
			inputT = utils.FindAllInts(line)
		} else {
			inputD = utils.FindAllInts(line)
		}
	})

	// Calculate for each race
	product := 1
	for i := range inputT {
		product *= possibleWins(inputT[i], inputD[i])
	}
	return product
}

func (Solution) Part2(input string) any {
	// Parse input
	var T, D int
	utils.ForEachLine(input, func(line string) {
		line = strings.ReplaceAll(line, " ", "")
		if T == 0 {
			T = utils.FindAllInts(line)[0]
		} else {
			D = utils.FindAllInts(line)[0]
		}
	})

	return possibleWins(T, D)
}
