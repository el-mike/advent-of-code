package day_12

import "fmt"

type GridCell struct {
	X    int
	Y    int
	Char rune
}

func NewGridCell(x, y int, char rune) *GridCell {
	return &GridCell{
		X:    x,
		Y:    y,
		Char: char,
	}
}

func (gc *GridCell) ID() string {
	return fmt.Sprintf("%d|%d", gc.X, gc.Y)
}

func (gc *GridCell) Same(candidate *GridCell) bool {
	return gc.X == candidate.X && gc.Y == candidate.Y
}

func (gc *GridCell) CanEnter(candidate *GridCell) bool {
	return gc.GetElevation() >= (candidate.GetElevation() - 1)
}

func (gc *GridCell) GetElevation() int {
	asciiCode := int(gc.Char)

	// For "S", we assume the lowest possible position.
	if asciiCode == 'S' {
		asciiCode = int('a')
	}

	// For "E", we assume the highest possible position.
	if asciiCode == 'E' {
		asciiCode = int('z')
	}
	return asciiCode - (AsciiLowercaseA - 1)
}
