package main

import (
	"2025/util"
	_ "embed"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

type QuantumBeam struct {
	Position  int
	Timelines int
}

//go:embed input.txt
var input string
var beams = []QuantumBeam{}

func main() {
	lines := strings.Split(input, "\n")
	// initial beam
	beams = append(beams, QuantumBeam{
		Position:  strings.Index(lines[0], "S"),
		Timelines: 1,
	})

	fmt.Printf("part 1: %d\n", part(1, lines[1:]))
	fmt.Printf("part 2: %d\n", part(2, lines[1:]))
}

func part(part int, lines []string) int {
	reSplitter := regexp.MustCompile(`\^`)
	if part == 1 {
		splits := 0
		for _, l := range lines {
			splitters := util.FindAllStartingIndex(reSplitter, []byte(l))
			for _, splitter := range splitters {
				// check if splitter is positioned in front of a beam
				for i, qb := range beams {
					if qb.Position == splitter {
						// delete unsplit beam
						beams = slices.Delete(beams, i, i+1)
						// split to left and right
						splitBeam(qb, -1)
						splitBeam(qb, 1)
						splits++
					}
				}
			}
		}
		return splits
	} else { // part 2
		timelines := 0
		for _, qb := range beams {
			timelines += qb.Timelines
		}
		return timelines
	}
}

// 1. check space for beam
// 2. if none, new beam with unsplit beam's timelines
// 3. if true, add unsplit beam's timelines to beam
func splitBeam(beam QuantumBeam, offset int) {
	beamOverlapIndex := -1
	for i, qb := range beams {
		if beam.Position+offset == qb.Position {
			beamOverlapIndex = i
			break
		}
	}
	if beamOverlapIndex != -1 {
		beams[beamOverlapIndex].Timelines += beam.Timelines
	} else {
		beams = append(beams, QuantumBeam{
			Position:  beam.Position + offset,
			Timelines: beam.Timelines,
		})
	}
}
