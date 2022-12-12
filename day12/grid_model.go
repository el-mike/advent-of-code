package day12

import "fmt"

const (
	AsciiLowercaseA = 97
)

type GridModel struct {
	Grid [][]*GridCell
	Rows int
	Cols int
	End  *GridCell
}

func NewGridModel(r, c int) *GridModel {
	return &GridModel{
		Grid: [][]*GridCell{},
		Rows: r,
		Cols: c,
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

func (g *GridModel) Parse(gridStr []string) {
	for x, line := range gridStr {
		g.Grid = append(g.Grid, []*GridCell{})

		for y, r := range line {
			cell := NewGridCell(x, y, r)

			g.Grid[x] = append(g.Grid[x], cell)

			if cell.Char == 'E' {
				g.End = cell
			}
		}
	}
}

func (g *GridModel) ForEach(cb func(cell *GridCell, x, y int)) {
	for x := range g.Grid {
		for y := range g.Grid[x] {
			cell := g.Grid[x][y]

			cb(cell, x, y)
		}
	}
}

func (g *GridModel) Render() {

	g.ForEach(func(cell *GridCell, x, y int) {
		fmt.Print(string(cell.Char))

		if y == (g.Cols - 1) {
			fmt.Println()
		}
	})
}
