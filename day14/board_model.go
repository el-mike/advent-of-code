package day14

import "fmt"

type Tile uint8
type Board [][]Tile

const (
	EmptyTile Tile = iota
	RockTile
	SandMotionTile
	SandRestTile
	StartTile
)

var tileSymbolMap = map[Tile]string{
	EmptyTile:      ".",
	RockTile:       "#",
	SandMotionTile: "~",
	SandRestTile:   "o",
	StartTile:      "+",
}

const BoardPadding = 10

type BoardModel struct {
	Board Board

	StartX int
	StartY int
	MaxX   int
	MaxY   int
}

func NewBoardModel(
	maxX,
	maxY int,
	startCoord *Coord,
	rockFormations [][]*Coord,
) *BoardModel {
	var board Board

	// We add 2 to accommodate for the "floor".
	maxY = maxY + 2

	for x := 0; x <= maxX; x += 1 {
		board = append(board, []Tile{})

		for y := 0; y <= maxY; y += 1 {
			if y == maxY {
				board[x] = append(board[x], RockTile)
			} else {
				board[x] = append(board[x], EmptyTile)
			}
		}
	}

	board[startCoord[0]][startCoord[1]] = StartTile

	for _, formation := range rockFormations {
		previous := formation[0]

		for i, current := range formation {
			if i == 0 {
				continue
			}

			startX := previous[0]
			endX := current[0]

			startY := previous[1]
			endY := current[1]

			if startX != endX {
				y := current[1]

				if startX > endX {
					startX, endX = endX, startX
				}

				for x := startX; x <= endX; x += 1 {
					board[x][y] = RockTile
				}
			}

			if startY != endY {
				x := current[0]

				if startY > endY {
					startY, endY = endY, startY
				}

				for y := startY; y <= endY; y += 1 {
					board[x][y] = RockTile
				}
			}

			previous = current

		}
	}

	return &BoardModel{
		Board:  board,
		StartX: startCoord[0],
		StartY: startCoord[1],
		MaxX:   maxX,
		MaxY:   maxY,
	}
}

func (bm *BoardModel) PadLeft() Board {
	bm.pad(true)

	return bm.Board
}

func (bm *BoardModel) PadRight() Board {
	bm.pad(false)

	return bm.Board
}

func (bm *BoardModel) pad(prepend bool) {
	var columns Board

	for x := 0; x < BoardPadding; x += 1 {
		columns = append(columns, []Tile{})

		for y := 0; y <= bm.MaxY; y += 1 {
			if y == bm.MaxY {
				columns[x] = append(columns[x], RockTile)
			} else {
				columns[x] = append(columns[x], EmptyTile)
			}
		}
	}

	if prepend {
		bm.Board = append(columns, bm.Board...)
		bm.StartX += BoardPadding
	} else {
		bm.Board = append(bm.Board, columns...)
		bm.MaxX += BoardPadding
	}
}

func (bm *BoardModel) Render(startX, startY int) {
	for y := startY; y <= bm.MaxY; y += 1 {
		for x := startX; x <= bm.MaxX; x += 1 {
			fmt.Print(tileSymbolMap[bm.Board[x][y]])
		}

		fmt.Println()
	}
}
