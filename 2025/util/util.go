package util

import (
	"math"
	"regexp"
)

type Position struct {
	X, Y int
}

type Position3D struct {
	X, Y, Z int
}

func AddPositions(a Position, b Position) Position {
	return Position{X: a.X + b.X, Y: a.Y + b.Y}
}

func AreEqualPositions(a Position, b Position) bool {
	return a.X == b.X && a.Y == b.Y
}

func GetSurroundingPositions() []Position {
	directions := []Position{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if j == 0 && i == 0 {
				continue
			}
			directions = append(directions, Position{
				X: i,
				Y: j,
			})
		}
	}
	return directions
}

func FindAllStartingIndex(re *regexp.Regexp, b []byte) []int {
	indexes := re.FindAllIndex(b, -1)
	var rs []int
	for _, ranges := range indexes {
		rs = append(rs, ranges[0])
	}
	return rs
}

func GetDistance3D(a Position3D, b Position3D) float64 {
	xDist := math.Pow((float64)(a.X-b.X), 2)
	yDist := math.Pow((float64)(a.Y-b.Y), 2)
	zDist := math.Pow((float64)(a.Z-b.Z), 2)
	return math.Sqrt(xDist + yDist + zDist)
}
