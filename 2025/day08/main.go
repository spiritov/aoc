package main

import (
	"2025/util"
	"cmp"
	_ "embed"
	"fmt"
	"log"
	"maps"
	"slices"
	"strings"
)

type Position3DPair struct {
	A, B util.Position3D
}

//go:embed input.txt
var input string

// each circuit is an array of junction boxes
var circuits = [][]util.Position3D{}
var distances = make(map[float64]Position3DPair)

func main() {
	lines := strings.Split(input, "\n")
	for i, l := range lines {
		circuits = append(circuits, []util.Position3D{})

		var x, y, z int
		fmt.Sscanf(l, "%d,%d,%d", &x, &y, &z)
		circuits[i] = append(circuits[i], util.Position3D{
			X: x,
			Y: y,
			Z: z,
		})
	}

	fmt.Printf("part 1: %d\n", part(1))
	fmt.Printf("part 2: %d\n", part(2))
}

func part(part int) int {
	if part == 1 {
		for i := 0; i <= len(circuits)-2; i++ {
			for j := 1; j < len(circuits)-i; j++ {
				currbox := circuits[i][0]
				nextBox := circuits[i+j][0]

				distances[util.GetDistance3D(currbox, nextBox)] = Position3DPair{
					A: currbox,
					B: nextBox,
				}
			}
		}

		// sort distances and make X number of connections
		sortedDistKeys := slices.Sorted(maps.Keys(distances))

		connections := 1000
		for i, distance := range sortedDistKeys {
			if i <= connections-1 {
				connectCircuits(distances[distance].A, distances[distance].B)
			}
		}

		// sort circuits desc by their length
		slices.SortFunc(circuits, func(a, b []util.Position3D) int {
			return cmp.Compare(len(b), len(a))
		})
		return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
	} else { // i'll make it right again
		sortedDistKeys := slices.Sorted(maps.Keys(distances))
		for _, distance := range sortedDistKeys {
			connectCircuits(distances[distance].A, distances[distance].B)

			// but it's no use, you said
			foundSoloCircuit := false
			for _, circuit := range circuits {
				// as my hunger grows and grows
				if len(circuit) == 1 {
					foundSoloCircuit = true
					break
				}
			}
			// i have to write the meaning of my life
			if !foundSoloCircuit {
				return distances[distance].A.X * distances[distance].B.X
			}
		}
	}
	return part
}

func connectCircuits(a util.Position3D, b util.Position3D) {
	aCircuitIndex := -1
	bCircuitIndex := -1
	for i, circuit := range circuits {
		if aCircuitIndex == -1 && slices.Contains(circuit, a) {
			aCircuitIndex = i
		}
		if bCircuitIndex == -1 && slices.Contains(circuit, b) {
			bCircuitIndex = i
		}
	}

	if aCircuitIndex == -1 || bCircuitIndex == -1 {
		log.Fatal("(っ◔◡◔)っ Matter has been lost.")
	}

	if aCircuitIndex == bCircuitIndex {
		return
	} else { // merge time
		circuits[aCircuitIndex] = append(circuits[aCircuitIndex], circuits[bCircuitIndex]...)
		circuits = slices.Delete(circuits, bCircuitIndex, bCircuitIndex+1)
	}
}
