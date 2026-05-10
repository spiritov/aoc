package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	Min int
	Max int
}

//go:embed input.txt
var input string
var ranges []Range

func main() {
	lines := strings.Split(input, "\n")

	fmt.Printf("part 1: %d\n", part(1, lines))
	fmt.Printf("part 2: %d\n", part(2, lines))
}

func part(part int, lines []string) int {
	ranges = nil
	freshIDs := 0
	scanIngredients := false
	for _, l := range lines {
		if len(l) == 0 {
			scanIngredients = true
		}
		if !scanIngredients {
			// populate ranges
			min, max := getMinMax(l)
			ranges = append(ranges, Range{
				Min: min,
				Max: max,
			})
			continue
		}
		if part == 1 {
			// testing ingredients
			for _, r := range ranges {
				iid, _ := strconv.Atoi(l)
				if iid >= r.Min && iid <= r.Max {
					freshIDs++
					break
				}
			}
		} else { // part 2
			// sort ranges by ascending mins
			slices.SortFunc(ranges, func(a, b Range) int {
				return cmp.Compare(a.Min, b.Min)
			})

			// check for prev range overlapping
			for i := range ranges[1:] {
				tryRangeMerge(i)
			}

			// add all IDs
			for _, r := range ranges {
				freshIDs += r.Max - r.Min + 1
			}
			break
		}
	}
	return freshIDs
}

func getMinMax(l string) (int, int) {
	var min int
	var max int
	fmt.Sscanf(l, "%d-%d", &min, &max)
	return min, max
}

func tryRangeMerge(i int) {
	if i >= len(ranges)-1 {
		return
	}
	prev := ranges[i]
	curr := ranges[i+1]
	if curr.Min <= prev.Max && curr.Max >= prev.Min {
		// combine ranges
		ranges[i].Min = min(prev.Min, curr.Min)
		ranges[i].Max = max(prev.Max, curr.Max)
		ranges = slices.Delete(ranges, i+1, i+2)
		tryRangeMerge(i)
	}
}
