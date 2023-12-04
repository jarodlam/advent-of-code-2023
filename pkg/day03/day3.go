package day03

import (
	"strconv"

	"github.com/jarodlam/advent-of-code-2023/pkg/utils"
)

type Solution struct{}

func iterateWindow3x3(iMid int, jMid int, grid [][]rune, f func(int, int, rune)) {
	iMax, jMax := utils.GridSize(grid)

	// Iterate over a 3x3 window
	for i := iMid - 1; i <= iMid+1; i++ {
		// Check if in bounds
		if i < 0 || i >= iMax {
			continue
		}

		for j := jMid - 1; j <= jMid+1; j++ {
			// Check if in bounds
			if j < 0 || j >= jMax {
				continue
			}

			// Run function
			f(i, j, grid[i][j])
		}
	}
}

func hasSymbol(iMid int, jMid int, grid [][]rune) (found bool) {
	iterateWindow3x3(iMid, jMid, grid, func(_ int, _ int, value rune) {
		// Skip if number or "."
		_, err := strconv.Atoi(string(value))
		if err == nil || value == '.' {
			return
		}
		found = true
	})
	return
}

func findGears(iMid int, jMid int, grid [][]rune, gearIdxs *map[[2]int]struct{}) {
	iterateWindow3x3(iMid, jMid, grid, func(i int, j int, value rune) {
		if value != '*' {
			return
		}
		(*gearIdxs)[[2]int{i, j}] = struct{}{}
	})
	return
}

func (Solution) Part1(input string) any {
	grid := utils.InputToGrid(input)

	total := 0
	numberStr := ""
	foundSymbol := false

	processNumber := func() {
		if numberStr != "" && foundSymbol {
			n, _ := strconv.Atoi(numberStr)
			total += n
		}
		numberStr = ""
		foundSymbol = false
	}

	for i, row := range grid {
		processNumber()

		for j := 0; j < len(row); j++ {
			value := row[j]

			// If number, append
			if _, err := strconv.Atoi(string(value)); err == nil {
				numberStr = numberStr + string(value)
				if hasSymbol(i, j, grid) {
					foundSymbol = true
				}
				continue
			}

			processNumber()
		}
	}

	return total
}

func (Solution) Part2(input string) any {
	grid := utils.InputToGrid(input)

	// Search for all gear symbols, keep track of numbers adjacent to them
	gearCounts := make(map[[2]int][]int)
	total := 0
	numberStr := ""
	currGearIdxs := make(map[[2]int]struct{})

	processNumber := func() {
		if numberStr != "" {
			n, _ := strconv.Atoi(numberStr)
			for idx := range currGearIdxs {
				gearCounts[[2]int{idx[0], idx[1]}] = append(gearCounts[[2]int{idx[0], idx[1]}], n)
			}
		}
		numberStr = ""
		currGearIdxs = make(map[[2]int]struct{})
	}

	for i, row := range grid {
		processNumber()

		for j := 0; j < len(row); j++ {
			value := row[j]

			// If number, append
			if _, err := strconv.Atoi(string(value)); err == nil {
				numberStr = numberStr + string(value)
				findGears(i, j, grid, &currGearIdxs)
				continue
			}

			processNumber()
		}
	}

	// Tally up all the valid gears
	for _, numbers := range gearCounts {
		if len(numbers) != 2 {
			continue
		}
		total += numbers[0] * numbers[1]
	}

	return total
}
