package day14

import (
	"strconv"
	"strings"
)

type Coord [2]int

func NewCoord(coordStr string) *Coord {
	parts := strings.Split(coordStr, ",")

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return &Coord{x, y}
}
