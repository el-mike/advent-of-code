package day12

import (
	"fmt"
	"github.com/fatih/color"
)

const (
	AsciiLowercaseA = 97
)

type GridModel struct {
	Grid [][]*GridCell
	Rows int
	Cols int
	End  *GridCell
}

func NewGridModel(gridStr []string) *GridModel {
	var grid [][]*GridCell
	var end *GridCell

	for x, line := range gridStr {
		grid = append(grid, []*GridCell{})

		for y, r := range line {
			cell := NewGridCell(x, y, r)

			grid[x] = append(grid[x], cell)

			if cell.Char == 'E' {
				end = cell
			}
		}
	}

	return &GridModel{
		Grid: grid,
		Rows: len(grid),
		Cols: len(grid[0]),
		End:  end,
	}
}

func (g *GridModel) GetAt(x, y int) *GridCell {
	return g.Grid[x][y]
}

func (g *GridModel) GetNeighbors(cell *GridCell) []*GridCell {
	var candidates []*GridCell

	if cell.X > 0 {
		candidates = append(candidates, g.GetAt(cell.X-1, cell.Y))
	}
	if cell.X < (g.Rows - 1) {
		candidates = append(candidates, g.GetAt(cell.X+1, cell.Y))
	}
	if cell.Y > 0 {
		candidates = append(candidates, g.GetAt(cell.X, cell.Y-1))
	}
	if cell.Y < (g.Cols - 1) {
		candidates = append(candidates, g.GetAt(cell.X, cell.Y+1))
	}

	var neighbors []*GridCell

	for _, candidate := range candidates {
		if cell.CanEnter(candidate) {
			neighbors = append(neighbors, candidate)
		}
	}

	return neighbors
}

func (g *GridModel) ForEach(cb func(cell *GridCell, x, y int)) {
	for x := range g.Grid {
		for y := range g.Grid[x] {
			cell := g.Grid[x][y]

			cb(cell, x, y)
		}
	}
}

func (g *GridModel) Render(path []*GridCell) {
	g.ForEach(func(cell *GridCell, x, y int) {
		isPath := false
		for _, pathCell := range path {
			if cell.Same(pathCell) {
				isPath = true
				break
			}
		}

		charStr := string(cell.Char)
		if isPath {
			c := color.New(color.FgGreen)
			c.Print(charStr)
		} else {
			fmt.Print(charStr)
		}

		if y == (g.Cols - 1) {
			fmt.Println()
		}
	})
}
