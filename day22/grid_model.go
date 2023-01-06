package day22

import (
	"el-mike/advent-of-code/common"
	"fmt"
)

type GridCell uint8
type CubeFace uint8

const (
	VoidCell GridCell = iota
	OpenCell
	WallCell
)

const (
	Face1 CubeFace = iota
	Face2
	Face3
	Face4
	Face5
	Face6
)

type GridModel struct {
	Grid           [][]GridCell
	Width          int
	Height         int
	FaceSize       int
	RowsEdgeCoords map[int]Coord
	ColsEdgeCoords map[int]Coord
}

func NewGridModel(faceSize int) *GridModel {
	return &GridModel{
		FaceSize: faceSize,
		Grid:     [][]GridCell{},
	}
}

func (gm *GridModel) Build(lines []string, maxRow int) {
	gm.Width = maxRow
	gm.Height = len(lines)

	for _, line := range lines {
		var gridRow []GridCell

		for i := 0; i < maxRow; i++ {
			var cell GridCell

			if i >= len(line) {
				cell = VoidCell
			} else {
				switch line[i] {
				case ' ':
					cell = VoidCell
				case '.':
					cell = OpenCell
				case '#':
					cell = WallCell
				}
			}

			gridRow = append(gridRow, cell)
		}

		gm.Grid = append(gm.Grid, gridRow)
	}
}

// GetTestFace - return the cube's face number given coord is on.
// Please note that this is specific, not general solution, and this
// function works for conrecte input. If input would be change, this function
// would need to bo changed as well to accommodate for different faces.
func (gm *GridModel) GetFace(coord Coord) CubeFace {
	faceSize := gm.FaceSize
	x, y := coord[0], coord[1]

	if y < faceSize {
		if x < 2*faceSize {
			return Face1
		}

		return Face2
	}

	if y < 2*faceSize {
		return Face3
	}

	if y < 3*faceSize {
		if x < faceSize {
			return Face4
		}

		return Face5
	}

	return Face6
}

func (gm *GridModel) GetWrapped(coord Coord, direction Direction) (Coord, Direction) {
	faceSize := gm.FaceSize
	currentFace := gm.GetFace(coord)

	xOffset, yOffset := coord[0]%faceSize, coord[1]%faceSize

	_, f1Bottom, f1Left, f1Top := gm.GetBorders(Face1)
	f2Right, f2Bottom, f2Left, f2Top := gm.GetBorders(Face2)
	f3Right, _, f3Left, f3Top := gm.GetBorders(Face3)
	_, f4Bottom, f4Left, f4Top := gm.GetBorders(Face4)
	f5Right, f5Bottom, f5Left, _ := gm.GetBorders(Face5)
	f6Right, f6Bottom, f6Left, f6Top := gm.GetBorders(Face6)

	if currentFace == Face1 {
		if direction == DirectionLeft {
			return Coord{f4Left, f4Bottom - yOffset}, DirectionRight
		}

		if direction == DirectionUp {
			return Coord{f6Left, f6Top + xOffset}, DirectionRight
		}
	}

	if currentFace == Face2 {
		if direction == DirectionUp {
			return Coord{f6Left + xOffset, f6Bottom}, DirectionUp
		}

		if direction == DirectionRight {
			return Coord{f5Right, f5Bottom - yOffset}, DirectionLeft
		}

		if direction == DirectionDown {
			return Coord{f3Right, f3Top + xOffset}, DirectionLeft
		}
	}

	if currentFace == Face3 {
		if direction == DirectionLeft {
			return Coord{f4Left + yOffset, f4Top}, DirectionDown
		}

		if direction == DirectionRight {
			return Coord{f2Left + yOffset, f2Bottom}, DirectionUp
		}
	}

	if currentFace == Face4 {
		if direction == DirectionLeft {
			return Coord{f1Left, f1Bottom - yOffset}, DirectionRight
		}

		if direction == DirectionUp {
			return Coord{f3Left, f3Top + xOffset}, DirectionRight
		}
	}

	if currentFace == Face5 {
		if direction == DirectionRight {
			return Coord{f2Right, f2Bottom - yOffset}, DirectionLeft
		}

		if direction == DirectionDown {
			return Coord{f6Right, f6Top + xOffset}, DirectionLeft
		}
	}

	if currentFace == Face6 {
		if direction == DirectionLeft {
			return Coord{f1Left + yOffset, f1Top}, DirectionDown
		}

		if direction == DirectionRight {
			return Coord{f5Left + yOffset, f5Bottom}, DirectionUp
		}

		if direction == DirectionDown {
			return Coord{f2Left + xOffset, f2Top}, DirectionDown
		}
	}

	return coord, direction
}

func (gm *GridModel) GetBorders(face CubeFace) (right, bottom, left, up int) {
	faceSize := gm.FaceSize

	if face == Face1 {
		return (2 * faceSize) - 1, faceSize - 1, faceSize, 0
	}

	if face == Face2 {
		return (3 * faceSize) - 1, faceSize - 1, 2 * faceSize, 0
	}

	if face == Face3 {
		return (2 * faceSize) - 1, (2 * faceSize) - 1, faceSize, faceSize
	}

	if face == Face4 {
		return faceSize - 1, (3 * faceSize) - 1, 0, 2 * faceSize
	}

	if face == Face5 {
		return (2 * faceSize) - 1, (3 * faceSize) - 1, faceSize, 2 * faceSize
	}

	return faceSize - 1, (4 * faceSize) - 1, 0, 3 * faceSize
}

func (gm *GridModel) Render(steps []Coord) {
	common.ClearTerminal()

	for y := 0; y < gm.Height; y++ {
		for x := 0; x < gm.Width; x++ {
			visited := false

			for _, step := range steps {
				if step[0] == x && step[1] == y {
					visited = true
					break
				}
			}

			if visited {
				fmt.Print("*")
			} else {
				switch gm.Grid[y][x] {
				case VoidCell:
					fmt.Print(" ")
				case OpenCell:
					fmt.Print(".")
				case WallCell:
					fmt.Print("#")
				}
			}
		}

		fmt.Println()
	}
}
