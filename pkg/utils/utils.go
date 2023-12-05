package utils

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ForEachLine maps function f over each line in input.
func ForEachLine(input string, f func(string)) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		f(line)
	}
}

// InputToGrid converts input text to a [][]rune grid.
func InputToGrid(input string) [][]rune {
	grid := make([][]rune, 0)
	ForEachLine(input, func(line string) {
		grid = append(grid, []rune(line))
	})
	return grid
}

// PrintGrid outputs a grid to the terminal.
func PrintGrid[T any](grid [][]T, space bool) {
	for _, line := range grid {
		for _, c := range line {
			switch x := any(c).(type) {
			case rune:
				fmt.Printf("%c", x)
			case bool:
				if x == true {
					fmt.Print("█")
				} else {
					fmt.Print(" ")
				}
			default:
				fmt.Print(x)
			}

			if space {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// GridSize retrieves the rows x cols dimensions of a grid.
func GridSize[T any](grid [][]T) (int, int) {
	return len(grid), len(grid[0])
}

// NewGrid makes a new empty grid.
func NewGrid[T any](rows int, cols int, c T) [][]T {
	grid := make([][]T, rows)
	for i := range grid {
		row := make([]T, cols)
		for j := range row {
			row[j] = c
		}
		grid[i] = row
	}
	return grid
}

// CloneGrid makes a new empty grid with the same dimensions as another grid.
func CloneGrid[T any, U any](gridToClone [][]U, c T) [][]T {
	rows, cols := GridSize(gridToClone)
	return NewGrid(rows, cols, c)
}

// FindAllInts extracts all integers from a string in order.
func FindAllInts(s string) []int {
	re := regexp.MustCompile(`\d+`)
	found := re.FindAllString(s, -1)
	numbers := make([]int, len(found))
	for i, x := range found {
		n, _ := strconv.Atoi(x)
		numbers[i] = n
	}
	return numbers
}
