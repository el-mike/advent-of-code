package day_17

type RockType uint8
type RockShape [][]bool

const (
	Horizontal RockType = iota
	Cross
	InvertedL
	Vertical
	Square
)

type Rock struct {
	Board    *Board
	RockType RockType
	Shape    RockShape
	Width    int
	Height   int
	Coords   Coords
	Settled  bool
}

var shapesMap = map[RockType]RockShape{
	Horizontal: {[]bool{true, true, true, true}},
	Cross: {
		[]bool{false, true, false},
		[]bool{true, true, true},
		[]bool{false, true, false},
	},
	InvertedL: {
		[]bool{false, false, true},
		[]bool{false, false, true},
		[]bool{true, true, true},
	},
	Vertical: {
		[]bool{true},
		[]bool{true},
		[]bool{true},
		[]bool{true},
	},
	Square: {
		[]bool{true, true},
		[]bool{true, true},
	},
}

func NewRock(
	board *Board,
	rockType RockType,
) *Rock {
	shape := shapesMap[rockType]

	return &Rock{
		Board:    board,
		RockType: rockType,
		Shape:    shape,
		Height:   len(shape),
		Width:    len(shape[0]),
		Settled:  false,
	}
}

func (r *Rock) Move(direction Direction) bool {
	diff := direction.GetTranslateValue()

	nextPosition := r.Coords.Clone()

	if direction.IsHorizontal() {
		nextPosition.TranslateX(diff)

		if !r.CheckCollision(nextPosition) {
			previousPosition := r.Coords
			r.Coords = nextPosition

			r.UpdatePosition(previousPosition)

			return false
		}
	}

	if direction.IsVertical() {
		nextPosition.TranslateY(diff)

		if r.CheckCollision(nextPosition) {
			r.Settled = true

			// When settled, we want to update board state as such for
			// given rock.
			r.Board.FillCells(r.Coords, SettledRockCell)

			rockMaxY := r.Coords.GetMaxY()

			if rockMaxY > r.Board.ReachedY {
				r.Board.ReachedY = rockMaxY
			}

			return true
		} else {
			previousPosition := r.Coords
			r.Coords = nextPosition

			r.UpdatePosition(previousPosition)
		}
	}

	return false
}

func (r *Rock) UpdatePosition(previousPosition Coords) {
	// First, we clear previously occupied cells.
	r.Board.FillCells(previousPosition, EmptyCell)
	r.Board.FillCells(r.Coords, FallingRockCell)
}

func (r *Rock) CheckCollision(nextPosition Coords) bool {
	for _, coord := range nextPosition {
		x, y := coord[0], coord[1]
		// If X coordinate is outside the board, it collides with the wall or floor.
		if x < 0 || x >= r.Board.Width || y < 0 {
			return true
		}

		targetCell := r.Board.Grid[y][x]

		// If next cell is not empty and not itself (falling rock), it means it collides
		// with something.
		if targetCell != EmptyCell && targetCell != FallingRockCell {
			return true
		}
	}

	return false
}
