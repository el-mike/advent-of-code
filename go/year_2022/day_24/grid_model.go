package day_24

import (
	"el-mike/advent-of-code/go/common"
	"fmt"
)

type GridCell uint8
type BlizzardPositions map[int]map[int][]Direction

const (
	OpenCell GridCell = iota
	BlizzardCell
	WallCell
)

const (
	OpenCellChar          = '.'
	WallCellChar          = '#'
	BlizzardCellRightChar = '>'
	BlizzardCellDownChar  = 'v'
	BlizzardCellLeftChar  = '<'
	BlizzardCellUpChar    = '^'
	ExpeditionCellChar    = 'E'
)

var BlizzardsMap = map[rune]Direction{
	BlizzardCellRightChar: DirectionRight,
	BlizzardCellDownChar:  DirectionDown,
	BlizzardCellLeftChar:  DirectionLeft,
	BlizzardCellUpChar:    DirectionUp,
}

type GridModel struct {
	BlizzardPositions BlizzardPositions
	Width             int
	Height            int
	Start             Coord
	End               Coord
}

func NewGridModel() *GridModel {
	return &GridModel{
		BlizzardPositions: BlizzardPositions{},
	}
}

func (gm *GridModel) Build(lines []string) {
	gm.Height = len(lines)
	gm.Width = len(lines[0])

	for y, line := range lines {
		gm.BlizzardPositions[y] = map[int][]Direction{}

		for x := 0; x < len(line); x++ {
			r := rune(line[x])

			// Get start coord.
			if y == 0 && r == OpenCellChar {
				gm.Start = NewCoord(x, y)
			}

			// Get end coord.
			if y == (gm.Height-1) && r == OpenCellChar {
				gm.End = NewCoord(x, y)
			}

			if r != OpenCellChar && r != WallCellChar {
				gm.BlizzardPositions[y][x] = []Direction{BlizzardsMap[r]}
			}
		}
	}
}

func (gm *GridModel) MoveBlizzards() {
	newPositions := BlizzardPositions{}

	for y, row := range gm.BlizzardPositions {

		for x, blizzards := range row {
			for _, blizzard := range blizzards {
				current := NewCoord(x, y)
				next := gm.GetNextBlizzard(current, blizzard)

				nextX, nextY := next.X, next.Y

				if _, ok := newPositions[nextY]; !ok {
					newPositions[nextY] = map[int][]Direction{}
				}

				newPositions[nextY][nextX] = append(newPositions[nextY][nextX], blizzard)
			}
		}
	}

	gm.BlizzardPositions = newPositions
}

func (gm *GridModel) GetNextBlizzard(coord Coord, direction Direction) Coord {
	if direction == DirectionRight {
		if coord.X == gm.Width-2 {
			return NewCoord(1, coord.Y)
		}

		return coord.GetRight()
	}

	if direction == DirectionLeft {
		if coord.X == 1 {
			return NewCoord(gm.Width-2, coord.Y)
		}

		return coord.GetLeft()
	}

	if direction == DirectionUp {
		if coord.Y == 1 {
			return NewCoord(coord.X, gm.Height-2)
		}

		return coord.GetUp()
	}

	if direction == DirectionDown {
		if coord.Y == gm.Height-2 {
			return NewCoord(coord.X, 1)
		}

		return coord.GetDown()
	}

	return coord
}

func (gm *GridModel) IsStart(coord Coord) bool {
	return coord.Same(gm.Start)
}

func (gm *GridModel) IsEnd(coord Coord) bool {
	return coord.Same(gm.End)
}

func (gm *GridModel) IsOpen(coord Coord) bool {
	return !gm.IsWall(coord) && !gm.hasBlizzard(coord)
}

func (gm *GridModel) IsWall(coord Coord) bool {
	if gm.IsStart(coord) || gm.IsEnd(coord) {
		return false
	}

	x, y := coord.X, coord.Y

	return (x == 0) || (x == gm.Width-1) || (y == 0) || (y == gm.Height-1)
}

func (gm *GridModel) hasBlizzard(coord Coord) bool {
	x, y := coord.X, coord.Y

	if gm.IsWall(coord) {
		return false
	}

	if row, ok := gm.BlizzardPositions[y]; ok {
		blizzards, ok := row[x]

		return ok && len(blizzards) > 0
	} else {
		return false
	}
}

func (gm *GridModel) Render(expeditionPosition Coord) {
	common.ClearTerminal()

	for y := 0; y < gm.Height; y++ {
		for x := 0; x < gm.Width; x++ {
			coord := NewCoord(x, y)

			if gm.hasBlizzard(coord) {
				blizzards := gm.BlizzardPositions[y][x]

				if len(blizzards) > 1 {
					fmt.Print(len(blizzards))
				} else {
					fmt.Print(string(common.GetKeyByValue(BlizzardsMap, blizzards[0])))
				}

			} else {
				if gm.IsOpen(coord) {
					if coord.Same(expeditionPosition) {
						fmt.Print(string(ExpeditionCellChar))
					} else {
						fmt.Print(string(OpenCellChar))
					}
				} else {
					fmt.Print(string(WallCellChar))
				}
			}
		}

		fmt.Println()
	}
}
