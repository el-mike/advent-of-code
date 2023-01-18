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
	Start             Vector
	End               Vector
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
				gm.Start = NewVector(x, y)
			}

			// Get end coord.
			if y == (gm.Height-1) && r == OpenCellChar {
				gm.End = NewVector(x, y)
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
				current := NewVector(x, y)
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

func (gm *GridModel) GetNextBlizzard(vector Vector, direction Direction) Vector {
	if direction == DirectionRight {
		if vector.X == gm.Width-2 {
			return NewVector(1, vector.Y)
		}

		return vector.GetRight()
	}

	if direction == DirectionLeft {
		if vector.X == 1 {
			return NewVector(gm.Width-2, vector.Y)
		}

		return vector.GetLeft()
	}

	if direction == DirectionUp {
		if vector.Y == 1 {
			return NewVector(vector.X, gm.Height-2)
		}

		return vector.GetUp()
	}

	if direction == DirectionDown {
		if vector.Y == gm.Height-2 {
			return NewVector(vector.X, 1)
		}

		return vector.GetDown()
	}

	return vector
}

func (gm *GridModel) IsStart(vector Vector) bool {
	return vector.Same(gm.Start)
}

func (gm *GridModel) IsEnd(vector Vector) bool {
	return vector.Same(gm.End)
}

func (gm *GridModel) IsOpen(vector Vector) bool {
	return !gm.IsWall(vector) && !gm.hasBlizzard(vector)
}

func (gm *GridModel) IsWall(vector Vector) bool {
	if gm.IsStart(vector) || gm.IsEnd(vector) {
		return false
	}

	x, y := vector.X, vector.Y

	return (x == 0) || (x == gm.Width-1) || (y == 0) || (y == gm.Height-1)
}

func (gm *GridModel) IsInBounds(vector Vector) bool {
	return vector.X >= 0 && vector.X < gm.Width && vector.Y >= 0 && vector.Y < gm.Height
}

func (gm *GridModel) hasBlizzard(coord Vector) bool {
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

func (gm *GridModel) Render(expeditionPosition Vector) {
	common.ClearTerminal()

	for y := 0; y < gm.Height; y++ {
		for x := 0; x < gm.Width; x++ {
			coord := NewVector(x, y)

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
