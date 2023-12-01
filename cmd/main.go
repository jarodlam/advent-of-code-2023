package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/jarodlam/advent-of-code-2023/pkg/common"
)

var (
	argDay       int
	argInputFile string
)

func solve(s pkg.Solution, input string) (any, any) {
	return s.Part1(input), s.Part2(input)
}

func defaultInputPath(day int) string {
	filename := fmt.Sprintf("day%02d.txt", day)
	return path.Join("input", filename)
}

func main() {
	// Arguments
	flag.IntVar(&argDay, "d", 0, "Day solution to run")
	flag.StringVar(&argInputFile, "i", "", "Input file to use, defaults to input/day<n>.txt")
	flag.Parse()

	// Get input
	if argInputFile == "" {
		argInputFile = defaultInputPath(argDay)
	}
	data, err := os.ReadFile(argInputFile)
	if err != nil {
		panic(err)
	}
	input := string(data)

	// Run solution
	sol := pkg.NewSolution(argDay)
	part1, part2 := solve(sol, input)
	fmt.Println(part1)
	fmt.Println(part2)
}
