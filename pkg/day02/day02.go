package day02

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/jarodlam/advent-of-code-2023/pkg/utils"
)

type Solution struct{}

type Draw struct {
	r int
	g int
	b int
}

type Game struct {
	number int
	draws  []Draw
}

func (game Game) IsPossible(r int, g int, b int) bool {
	for _, d := range game.draws {
		if d.r > r || d.g > g || d.b > b {
			return false
		}
	}
	return true
}

func (game Game) MinimumPower() int {
	rMax, gMax, bMax := 0, 0, 0
	for _, d := range game.draws {
		if d.r > rMax {
			rMax = d.r
		}
		if d.g > gMax {
			gMax = d.g
		}
		if d.b > bMax {
			bMax = d.b
		}
	}
	return rMax * gMax * bMax
}

func parseGame(line string) Game {
	// Game number
	gameStr, drawsStr, _ := strings.Cut(line, ": ")
	gameNo, _ := strconv.Atoi(gameStr[5:])

	// Draws
	drawsStrs := strings.Split(drawsStr, "; ")
	draws := []Draw{}
	for _, d := range drawsStrs {
		colourStrs := strings.Split(d, ", ")
		draw := Draw{}
		for _, c := range colourStrs {
			numColourStr, colour, _ := strings.Cut(c, " ")
			numColour, _ := strconv.Atoi(numColourStr)
			switch colour {
			case "red":
				draw.r = numColour
			case "green":
				draw.g = numColour
			case "blue":
				draw.b = numColour
			}
		}
		draws = append(draws, draw)
	}

	return Game{gameNo, draws}
}

func (Solution) Part1(input string) any {
	// Parse input
	// scanner := bufio.NewScanner(strings.NewReader(input))
	// games := []Game{}
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	games = append(games, parseGame(line))
	// }
	games := []Game{}
	utils.ForEachLine(input, func(line string) {
		games = append(games, parseGame(line))
	})

	// Tally possible games
	total := 0
	for _, game := range games {
		if game.IsPossible(12, 13, 14) {
			total += game.number
		}
	}

	return total
}

func (Solution) Part2(input string) any {
	// Parse input
	scanner := bufio.NewScanner(strings.NewReader(input))
	games := []Game{}
	for scanner.Scan() {
		line := scanner.Text()
		games = append(games, parseGame(line))
	}

	// Calculate powers
	total := 0
	for _, game := range games {
		total += game.MinimumPower()
	}

	return total
}
