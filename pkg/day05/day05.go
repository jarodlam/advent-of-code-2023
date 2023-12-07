package day05

import (
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

func (m mapping) forwardMapInterval(i interval) (interval, interval, bool) {
	mFrom := m.source
	mTo := m.source + m.length - 1
	if i.to < mFrom || mTo < i.from {
		// No overlap
		return interval{}, interval{}, false
	} else if i.from < mFrom && mFrom <= i.to && i.to <= mTo {
		// Right side overlap
		// fmt.Printf("        Right side overlap with %v\n", m)
		return interval{m.forwardMapNoCheck(mFrom), m.forwardMapNoCheck(i.to)}, interval{mFrom, i.to}, true
	} else if mFrom <= i.from && i.from <= mTo && mTo < i.to {
		// Left side overlap
		// fmt.Printf("        Left side overlap with %v\n", m)
		return interval{m.forwardMapNoCheck(i.from), m.forwardMapNoCheck(mTo)}, interval{i.from, mTo}, true
	} else if i.from < mFrom && mTo < i.to {
		// Middle overlap
		// fmt.Printf("        Middle overlap with %v\n", m)
		return interval{m.forwardMapNoCheck(mFrom), m.forwardMapNoCheck(mTo)}, interval{i.from, i.to}, true
	} else {
		// Full overlap
		// fmt.Printf("        Full overlap with %v\n", m)
		return interval{m.forwardMapNoCheck(i.from), m.forwardMapNoCheck(i.to)}, interval{i.from, i.to}, true
	}
}

func mergeIntervals(intvls []interval) []interval {
	if len(intvls) == 0 {
		return intvls
	}

	// Sort by from value
	sort.Slice(
		intvls,
		func(i, j int) bool { return intvls[i].from < intvls[j].from },
	)

	// Merge
	out := make([]interval, 0)
	prevInt := intvls[0]
	for i := 1; i < len(intvls); i++ {
		// If overlaps with prevInt, merge with prevInt
		if intvls[i].from <= prevInt.to {
			if prevInt.to < intvls[i].to {
				prevInt.to = intvls[i].to
			}
			continue
		}

		// No overlap, append prevInt and start working on a new int
		out = append(out, prevInt)
		prevInt = intvls[i]
	}
	out = append(out, prevInt)

	return out
}

func intervalDiff(a []interval, b interval) []interval {
	out := make([]interval, 0)
	for _, i := range a {
		if i.to < b.from || b.to < i.from {
			// No overlap, entire interval survives
			out = append(out, i)
		} else if i.from < b.from && b.from <= i.to && i.to <= b.to {
			// Right side overlap
			out = append(out, interval{i.from, b.from - 1})
		} else if b.from <= i.from && i.from <= b.to && b.to < i.to {
			// Left side overlap
			out = append(out, interval{b.to + 1, i.to})
		} else if i.from < b.from && b.to < i.to {
			// Middle overlap
			out = append(out, interval{i.from, b.from - 1}, interval{b.to + 1, i.to})
		} else {
			// Full overlap, entire interval subtracted
		}
	}
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
	// fmt.Printf("Processing seed interval %v\n", seed)
	seedIntervals := []interval{seed}
	for _, ms := range mappingSets {
		// fmt.Printf("  Stage %d\n", i)
		newSeedIntervals := make([]interval, 0)

		// Individually map each interval for this stage
		for _, intvl := range seedIntervals {
			// fmt.Printf("    Processing interval %v\n", intvl)
			newIntvl := make([]interval, 0)
			unmapped := []interval{intvl}
			// Map as much as we can
			for _, m := range ms {
				newValue, mappedPortion, ok := m.forwardMapInterval(intvl)
				if ok {
					newIntvl = append(newIntvl, newValue)
					unmapped = intervalDiff(unmapped, mappedPortion)
				}
			}

			// Merge and infill the rest
			// fmt.Printf("      Mapped to %v\n", newIntvl)
			newIntvl = mergeIntervals(newIntvl)
			// fmt.Printf("      Merged to %v\n", newIntvl)
			// newIntvl = infillIntervals(newIntvl, intvl)
			newIntvl = append(newIntvl, unmapped...)
			// fmt.Printf("      Infilled to %v\n", newIntvl)
			newSeedIntervals = append(newSeedIntervals, newIntvl...)
		}

		// Merge all individually mapped intervals
		seedIntervals = mergeIntervals(newSeedIntervals)
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

/*
Part2
Stepping through example solution
Seeds: 79
*/
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
				seeds = append(seeds, interval{seedNumbers[i], seedNumbers[i] + seedNumbers[i+1] - 1})
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
