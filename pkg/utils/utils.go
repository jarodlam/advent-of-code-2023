package utils

import (
	"bufio"
	"strings"
)

// Map function f over each line in input.
func ForEachLine(input string, f func(string)) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		f(line)
	}
}
