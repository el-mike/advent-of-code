package day22

import (
	"el-mike/advent-of-code/common"
	"fmt"
)

type GridCell uint8

const (
	VoidCell GridCell = iota
	OpenCell
	WallCell
)

type GridModel struct {
	Grid           [][]GridCell
	Width          int
	Height         int
	RowsEdgeCoords map[int]Coord
	ColsEdgeCoords map[int]Coord
}

func NewGridModel() *GridModel {
	return &GridModel{
		Grid: [][]GridCell{},
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

func (gm *GridModel) GetRowEdgeCells(y int) (Coord, Coord) {
	var leftEdge, rightEdge Coord

	for x := 0; x < gm.Width; x++ {
		cell := gm.Grid[y][x]

		if cell != VoidCell && leftEdge == NilCoord() {
			leftEdge = Coord{x, y}
		}

		if x == (gm.Width - 1) {
			rightEdge = Coord{x, y}
			break
		}

		if leftEdge != NilCoord() && gm.Grid[y][x+1] == VoidCell && rightEdge == NilCoord() {
			rightEdge = Coord{x, y}
			break
		}
	}

	return leftEdge, rightEdge
}

func (gm *GridModel) GetColEdgeCells(x int) (Coord, Coord) {
	var topEdge, bottomEdge Coord

	for y := 0; y < gm.Height; y++ {
		cell := gm.Grid[y][x]

		if cell != VoidCell && topEdge == NilCoord() {
			topEdge = Coord{x, y}
		}

		if y == (gm.Height - 1) {
			bottomEdge = Coord{x, y}
			break
		}

		if topEdge != NilCoord() && gm.Grid[y+1][x] == VoidCell && bottomEdge == NilCoord() {
			bottomEdge = Coord{x, y}
			break
		}
	}

	return topEdge, bottomEdge
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
