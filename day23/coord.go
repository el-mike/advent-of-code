package day23

import (
	"fmt"
	"strconv"
	"strings"
)

type Coord [2]int
type CoordsMap map[string]Coord

func (c Coord) GetKey() string {
	return fmt.Sprintf("%d|%d", c[0], c[1])
}

func NewCoordFromKey(key string) Coord {
	parts := strings.Split(key, "|")
	xStr, yStr := parts[0], parts[1]

	x, err := strconv.Atoi(xStr)
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(yStr)
	if err != nil {
		panic(err)
	}

	return Coord{x, y}
}

func (c Coord) GetNW() Coord {
	return Coord{c[0] - 1, c[1] - 1}
}

func (c Coord) GetN() Coord {
	return Coord{c[0], c[1] - 1}
}

func (c Coord) GetNE() Coord {
	return Coord{c[0] + 1, c[1] - 1}
}

func (c Coord) GetE() Coord {
	return Coord{c[0] + 1, c[1]}
}

func (c Coord) GetSE() Coord {
	return Coord{c[0] + 1, c[1] + 1}
}

func (c Coord) GetS() Coord {
	return Coord{c[0], c[1] + 1}
}

func (c Coord) GetSW() Coord {
	return Coord{c[0] - 1, c[1] + 1}
}

func (c Coord) GetW() Coord {
	return Coord{c[0] - 1, c[1]}
}

func (c Coord) GetByDirection(direction Direction) Coord {
	if direction == DirectionN {
		return c.GetN()
	}

	if direction == DirectionNE {
		return c.GetNE()
	}

	if direction == DirectionE {
		return c.GetE()
	}

	if direction == DirectionSE {
		return c.GetSE()
	}

	if direction == DirectionS {
		return c.GetS()
	}

	if direction == DirectionSW {
		return c.GetSW()
	}

	if direction == DirectionW {
		return c.GetW()
	}

	if direction == DirectionNW {
		return c.GetNW()
	}

	return c
}

func NilCoord() Coord {
	var nilCoord Coord

	return nilCoord
}
