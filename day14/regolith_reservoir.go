package day14

import (
	"el-mike/advent-of-code/common"
	"fmt"
	"strings"
	"time"
)

const PointDelimiter = " -> "

const RenderSimulation = false

func RegolithReservoir() {
	scanner, err := common.GetFileScanner("./day14/" + common.InputFilename)
	if err != nil {
		panic(err)
	}

	var rockFormations [][]*Coord

	// Please note that minY will stay 0, as the starting index of the sand particle
	// is at index 0.
	maxX, maxY := 0, 0
	for i := 0; scanner.Scan(); i += 1 {
		line := scanner.Text()

		if line == "" {
			continue
		}

		coordStrings := strings.Split(line, PointDelimiter)

		var rockFormation []*Coord

		for _, coordStr := range coordStrings {
			coord := NewCoord(coordStr)
			rockFormation = append(rockFormation, coord)

			if coord[0] > maxX {
				maxX = coord[0]
			}

			if coord[1] > maxY {
				maxY = coord[1]
			}
		}

		rockFormations = append(rockFormations, rockFormation)
	}

	startCoord := &Coord{500, 0}

	boardModel := NewBoardModel(maxX, maxY, startCoord, rockFormations)

	unitsCount := simulate(boardModel)

	fmt.Println(unitsCount)
}

func simulate(boardModel *BoardModel) int {
	sx, sy := boardModel.StartX, boardModel.StartY

	board := boardModel.Board

	unitsCount := 0

	for {
		board[sx][sy] = SandMotionTile

		if RenderSimulation {
			time.Sleep(20 * time.Millisecond)

			common.ClearTerminal()
			boardModel.Render(490, 0)
		}

		if sx == 1 {
			board = boardModel.PadLeft()
		}
		if sx == (boardModel.MaxX - 1) {
			board = boardModel.PadRight()
		}

		if board[sx][sy+1] == EmptyTile {
			board[sx][sy] = EmptyTile

			sy = sy + 1
		} else if board[sx-1][sy+1] == EmptyTile {
			board[sx][sy] = EmptyTile

			sx, sy = sx-1, sy+1
		} else if board[sx+1][sy+1] == EmptyTile {
			board[sx][sy] = EmptyTile

			sx, sy = sx+1, sy+1
		} else {
			board[sx][sy] = SandRestTile
			unitsCount += 1

			if sx == boardModel.StartX && sy == boardModel.StartY {
				break
			}

			sx, sy = boardModel.StartX, boardModel.StartY
		}
	}

	return unitsCount
}
