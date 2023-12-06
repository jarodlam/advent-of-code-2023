package day05

import (
	"fmt"
	"math"
	"sort"

	"github.com/jarodlam/advent-of-code-2023/pkg/utils"
)

type Solution struct{}

type mapping struct {
	dest   int
	source int
	length int
}

type mappingSet = []mapping

type interval struct {
	from int // inclusive
	to   int // inclusive
}

func (m mapping) forwardMap(value int) (int, bool) {
	// Check if value within range
	if !(m.source <= value && value < (m.source+m.length)) {
		return -1, false
	}

	return m.forwardMapNoCheck(value), true
}

func (m mapping) forwardMapNoCheck(value int) int {
	return m.dest + (value - m.source)
}

func (m mapping) forwardMapInterval(i interval) []interval {
	mFrom := m.source
	mTo := m.source + m.length - 1
	if i.from < mFrom && mFrom <= i.to && i.to <= mTo {
		// Right side overlap
		return []interval{
			{i.from, mFrom - 1},
			{m.forwardMapNoCheck(mFrom), m.forwardMapNoCheck(i.to)},
		}
	} else if mFrom <= i.from && i.from <= mTo && mTo < i.to {
		// Left side overlap
		return []interval{
			{m.forwardMapNoCheck(i.from), m.forwardMapNoCheck(mTo)},
			{mTo + 1, i.to},
		}
	} else if mFrom < i.from && i.to < mTo {
		// Middle overlap
		return []interval{
			{mFrom, i.from - 1},
			{m.forwardMapNoCheck(i.from), m.forwardMapNoCheck(i.to)},
			{i.to + 1, mTo},
		}
	} else {
		// Full overlap
		return []interval{
			{m.forwardMapNoCheck(i.from), m.forwardMapNoCheck(i.to)},
		}
	}
}

func mergeIntervals(ints []interval) []interval {
	// Sort by from value
	sort.Slice(
		ints,
		func(i, j int) bool { return ints[i].from < ints[j].from },
	)

	// Merge
	out := make([]interval, 0)
	prevInt := ints[0]
	for i := 1; i < len(ints); i++ {
		// If overlaps with prevInt, merge with prevInt
		if ints[i].from <= prevInt.to {
			if prevInt.to < ints[i].to {
				prevInt.to = ints[i].to
			}
			continue
		}

		// No overlap, append prevInt and start working on a new int
		out = append(out, prevInt)
		prevInt = ints[i]
	}
	out = append(out, prevInt)

	return out
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

func locateMinFromSeedRange(mappingSets []mappingSet, seed interval) int {
	fmt.Printf("Processing seed interval %v\n", seed)
	seedIntervals := []interval{seed}
	for _, ms := range mappingSets {
		newSeedIntervals := make([]interval, 0)
		// Check all available mappings
		for _, m := range ms {
			for _, intvl := range seedIntervals {
				newSeedIntervals = append(newSeedIntervals, m.forwardMapInterval(intvl)...)
			}
			fmt.Printf("  Mapped to %v\n", newSeedIntervals)
			seedIntervals = mergeIntervals(newSeedIntervals)
			fmt.Printf("    Merged to %v\n", seedIntervals)
		}
	}

	// Find min value
	minLocation := math.MaxInt
	for _, intvl := range seedIntervals {
		if intvl.from < minLocation {
			minLocation = intvl.from
		}
	}
	return minLocation
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
	// Parse
	maps := make([]mappingSet, 7)
	i := -1
	firstLine := true
	var seeds []interval
	utils.ForEachLine(input, func(line string) {
		// First line, seed numbers
		if firstLine {
			seedNumbers := utils.FindAllInts(line)
			for i := 0; i < len(seedNumbers); i += 2 {
				seeds = append(seeds, interval{seedNumbers[i], seedNumbers[i] + seedNumbers[i+1]})
			}
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

	fmt.Println(seeds)

	// Locations
	minLocation := math.MaxInt
	for _, seed := range seeds {
		loc := locateMinFromSeedRange(maps, seed)
		if loc < minLocation {
			minLocation = loc
		}
	}

	return minLocation
}
