package day_23

import (
	"el-mike/advent-of-code/go/common"
	"errors"
	"fmt"
)

type GridCell uint8

const (
	OpenCell GridCell = iota
	ElfCell
)

const (
	OpenCellChar = '.'
	ElfCellChar  = '#'
)

var CellCharMap = map[rune]GridCell{
	OpenCellChar: OpenCell,
	ElfCellChar:  ElfCell,
}

type GridModel struct {
	ElfPositions CoordsMap
	Width        int
	Height       int
}

func NewGridModel() *GridModel {
	return &GridModel{
		ElfPositions: CoordsMap{},
	}
}

func (gm *GridModel) Build(lines []string) {
	gm.Height = len(lines)
	gm.Width = len(lines[0])

	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			if CellCharMap[rune(line[x])] == ElfCell {
				coord := Coord{x, y}

				gm.ElfPositions[coord.GetKey()] = coord
			}
		}
	}
}

func (gm *GridModel) GetAt(coord Coord) GridCell {
	_, ok := gm.ElfPositions[coord.GetKey()]

	if ok {
		return ElfCell
	}

	return OpenCell
}

func (gm *GridModel) HasNeighborByDirection(coord Coord, direction Direction) (bool, error) {
	var n1, n2, n3 Coord

	if direction == DirectionN {
		n1, n2, n3 = coord.GetN(), coord.GetNW(), coord.GetNE()
	} else if direction == DirectionS {
		n1, n2, n3 = coord.GetS(), coord.GetSW(), coord.GetSE()
	} else if direction == DirectionW {
		n1, n2, n3 = coord.GetW(), coord.GetNW(), coord.GetSW()
	} else if direction == DirectionE {
		n1, n2, n3 = coord.GetE(), coord.GetNE(), coord.GetSE()
	} else {
		return false, errors.New("incorrect direction")
	}

	hasNeighbor := (!gm.ExtendsGrid(n1) && gm.GetAt(n1) == ElfCell) ||
		(!gm.ExtendsGrid(n2) && gm.GetAt(n2) == ElfCell) ||
		(!gm.ExtendsGrid(n3) && gm.GetAt(n3) == ElfCell)

	return hasNeighbor, nil
}

func (gm *GridModel) HasAnyNeighbor(coord Coord) bool {
	hasN, _ := gm.HasNeighborByDirection(coord, DirectionN)
	hasE, _ := gm.HasNeighborByDirection(coord, DirectionE)
	hasS, _ := gm.HasNeighborByDirection(coord, DirectionS)
	hasW, _ := gm.HasNeighborByDirection(coord, DirectionW)

	return hasN || hasE || hasS || hasW
}

func (gm *GridModel) ExtendsTop(coord Coord) bool {
	return coord[1] < 0
}

func (gm *GridModel) ExtendsRight(coord Coord) bool {
	return coord[0] >= gm.Width
}

func (gm *GridModel) ExtendsBottom(coord Coord) bool {
	return coord[1] >= gm.Height
}

func (gm *GridModel) ExtendsLeft(coord Coord) bool {
	return coord[0] < 0
}

func (gm *GridModel) ExtendsGrid(coord Coord) bool {
	return gm.ExtendsTop(coord) ||
		gm.ExtendsRight(coord) ||
		gm.ExtendsBottom(coord) ||
		gm.ExtendsLeft(coord)
}

func (gm *GridModel) PadIfExtends(padTop, padRight, padBottom, padLeft bool) {
	if padTop {
		gm.Height += 1

		gm.MoveAll(DirectionS)
	}

	if padRight {
		gm.Width += 1
	}

	if padBottom {
		gm.Height += 1
	}

	if padLeft {
		gm.Width += 1

		gm.MoveAll(DirectionE)
	}
}

func (gm *GridModel) MoveAll(direction Direction) {
	newPositions := CoordsMap{}
	for _, position := range gm.ElfPositions {
		var newPosition Coord

		if direction == DirectionS {
			newPosition = Coord{position[0], position[1] + 1}
		}

		if direction == DirectionE {
			newPosition = Coord{position[0] + 1, position[1]}
		}

		newPositions[newPosition.GetKey()] = newPosition
	}

	gm.ElfPositions = newPositions
}

func (gm *GridModel) Render() {
	common.ClearTerminal()

	for y := 0; y < gm.Height; y++ {
		for x := 0; x < gm.Width; x++ {
			if _, ok := gm.ElfPositions[Coord{x, y}.GetKey()]; ok {
				fmt.Print(string(ElfCellChar))
			} else {
				fmt.Print(string(OpenCellChar))
			}
		}

		fmt.Println()
	}
}
