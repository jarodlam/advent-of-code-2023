package day05

import (
	"math"

	"github.com/jarodlam/advent-of-code-2023/pkg/utils"
)

type Solution struct{}

type mapping struct {
	dest   int
	source int
	length int
}

type mappingSet = []mapping

func (m mapping) forwardMap(value int) (int, bool) {
	// Check if value within range
	if !(m.source <= value && value < (m.source+m.length)) {
		return -1, false
	}

	newValue := m.dest + (value - m.source)
	return newValue, true
}

func locateSeed(mappingSets []mappingSet, seed int) int {
	value := seed
	for _, ms := range mappingSets {
		// Check all available mappings
		// If not found, will stay on same value as specified
		for _, m := range ms {
			newValue, ok := m.forwardMap(value)
			if ok {
				value = newValue
				break
			}
		}
	}
	return value
}

func (Solution) Part1(input string) any {
	// Parse
	maps := make([]mappingSet, 7)
	i := -1
	firstLine := true
	var seeds []int
	utils.ForEachLine(input, func(line string) {
		// First line, seed numbers
		if firstLine {
			seeds = utils.FindAllInts(line)
			firstLine = false
			return
		}

		// Blank line is separator between lists
		if line == "" {
			i += 1
			maps[i] = make(mappingSet, 0)
			return
		}

		numbers := utils.FindAllInts(line)

		// Ignore section heading
		if len(numbers) == 0 {
			return
		}

		// Append
		maps[i] = append(maps[i], mapping{numbers[0], numbers[1], numbers[2]})
	})

	// Locations
	minLocation := math.MaxInt
	for _, seed := range seeds {
		loc := locateSeed(maps, seed)
		if loc < minLocation {
			minLocation = loc
		}
	}

	return minLocation
}

func (Solution) Part2(input string) any {
	return 0
}
