package main

import (
	"2025/util"
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string
var positions []util.Position
var deletables []util.Position
var directions = util.GetSurroundingPositions()

func main() {
	lines := strings.Split(input, "\n")
	positions = getPaperPositions(lines)

	fmt.Printf("part 1: %d\n", part(1))
	fmt.Printf("part 2: %d\n", part(2))
}

func part(part int) int {
	sum := 0

	if part == 1 {
		sum = getSumOfRemovables(1)
	} else {
		for {
			loopSum := getSumOfRemovables(2)
			sum += loopSum
			if (loopSum) == 0 {
				break
			}

			positions = slices.DeleteFunc(positions, func(p util.Position) bool {
				return slices.Contains(deletables, p)
			})
		}
	}

	return sum
}

func getPaperPositions(lines []string) []util.Position {
	var positions []util.Position
	for i, line := range lines {
		for j, obj := range line {
			if obj == '@' {
				positions = append(positions, util.Position{X: i, Y: j})
			}
		}
	}
	return positions
}

func getSumOfRemovables(part int) int {
	sum := 0

	for _, position := range positions {
		neighbors := 0

		for _, direction := range directions {
			if slices.Contains(positions, util.AddPositions(position, direction)) {
				neighbors++
			}
		}

		if neighbors < 4 {
			sum++
			if part == 2 {
				deletables = append(deletables, position)
			}
		}
	}

	return sum
}
