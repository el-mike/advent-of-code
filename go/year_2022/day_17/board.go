package day_17

import (
	"fmt"
	"strings"
)

type BoardCell uint8

const (
	EmptyCell BoardCell = iota
	FallingRockCell
	SettledRockCell
)

const (
	BoardWidth           = 7
	NewRockBottomPadding = 3
	NewRockLeftPadding   = 2
	RenderHeight         = 10
)

type Board struct {
	Width int

	Grid          [][]BoardCell
	FallingRock   *Rock
	MaxY          int
	ReachedY      int
	InitialReachY int
}

var cellSignsMap = map[BoardCell]string{
	EmptyCell:       ".",
	FallingRockCell: "@",
	SettledRockCell: "#",
}

func NewBoard(initialState [][]BoardCell) *Board {
	var grid [][]BoardCell

	// We need to start at -1, as ReachedY means concrete index,
	// and when starting, no index has been reached yet.
	reachedY := -1
	initialReachedY := -1

	if initialState != nil {
		grid = initialState

		// If we start with some defined levels, we want to set
		// reachedY accordingly.
		initialReachedY = len(initialState) - 1
		reachedY = initialReachedY
	}

	return &Board{
		Width:         BoardWidth,
		Grid:          grid,
		MaxY:          0,
		ReachedY:      reachedY,
		InitialReachY: initialReachedY,
	}
}

func (b *Board) AddRock(rock *Rock) {
	numRows := NewRockBottomPadding + rock.Height

	// We need to subtract one to make MaxY an actual highest row index,
	// not the height (length) itself.
	b.MaxY = b.ReachedY + numRows

	for y := b.ReachedY; y <= b.MaxY; y += 1 {
		b.AddRow(y)
	}

	var coords Coords

	for y, row := range rock.Shape {
		gridY := b.MaxY - y

		for x, cell := range row {
			gridX := NewRockLeftPadding + x

			// For every actual rock cell (not empty space from shape),
			// we add a new coord for given rock.
			if cell {
				coords = append(coords, Coord{gridX, gridY})

				b.Grid[gridY][gridX] = FallingRockCell
			}
		}
	}

	rock.Coords = coords

	b.FallingRock = rock
}

func (b *Board) FillCells(coords Coords, cell BoardCell) {
	for _, coord := range coords {
		x, y := coord[0], coord[1]

		b.Grid[y][x] = cell
	}
}

func (b *Board) AddRow(y int) {
	if y < len(b.Grid) {
		return
	}

	b.Grid = append(b.Grid, []BoardCell{})

	for x := 0; x < b.Width; x += 1 {
		b.Grid[y] = append(b.Grid[y], EmptyCell)
	}
}

func (b *Board) GetTopRow() []BoardCell {
	return b.Grid[b.ReachedY]
}

func (b *Board) GetTopRows(numRows int) [][]BoardCell {
	var result [][]BoardCell

	delimiter := b.ReachedY - numRows
	if delimiter < 0 {
		delimiter = 0
	}

	for i := b.ReachedY; i >= delimiter; i -= 1 {
		result = append(result, b.Grid[i])
	}

	return result
}

func (b *Board) Render() {
	delimiter := b.MaxY - RenderHeight

	if delimiter < 0 {
		delimiter = 0
	}

	fmt.Printf("\n")

	for y := b.MaxY; y >= delimiter; y -= 1 {
		fmt.Print("|")

		for _, cell := range b.Grid[y] {
			fmt.Print(cellSignsMap[cell])
		}

		fmt.Println("|")
	}

	bottomBorder := strings.Repeat("-", b.Width)

	fmt.Println("+" + bottomBorder + "+")
	fmt.Printf("\n\n")
}
