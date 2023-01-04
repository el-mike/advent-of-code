package day22

import (
	"el-mike/advent-of-code/common"
	"fmt"
)

func MonkeyMap() {
	scanner, err := common.GetFileScanner("./day22/" + common.InputFilename)
	if err != nil {
		panic(err)
	}

	lastRowReached := false

	var lines []string
	var movesStr string

	maxRow := 0

	for scanner.Scan() {
		line := scanner.Text()

		if lastRowReached {
			movesStr = line
			break
		}

		if line == "" {
			lastRowReached = true
			continue
		}

		lines = append(lines, line)

		if len(line) > maxRow {
			maxRow = len(line)
		}
	}

	gridModel := NewGridModel()
	instructions := NewInstructions(movesStr)

	gridModel.Build(lines, maxRow)

	currentCoord := Coord{0, 0}

	for x, cell := range gridModel.Grid[0] {
		if cell != VoidCell {
			currentCoord[0] = x
			break
		}
	}

	steps := []Coord{currentCoord}

	currentDirection := DirectionRight

	for i := 0; i < len(instructions.Numbers); i++ {
		n := instructions.Numbers[i]

		for j := 0; j < n; j++ {
			var nextCoord Coord

			switch currentDirection {
			case DirectionLeft:
				nextCoord = Coord{currentCoord[0] - 1, currentCoord[1]}
			case DirectionRight:
				nextCoord = Coord{currentCoord[0] + 1, currentCoord[1]}
			case DirectionUp:
				nextCoord = Coord{currentCoord[0], currentCoord[1] - 1}
			case DirectionDown:
				nextCoord = Coord{currentCoord[0], currentCoord[1] + 1}
			}

			nextX, nextY := nextCoord[0], nextCoord[1]

			var nextCell GridCell

			if nextX < 0 || nextX >= gridModel.Width || nextY < 0 || nextY >= gridModel.Height {
				nextCell = VoidCell
			} else {
				nextCell = gridModel.Grid[nextY][nextX]
			}

			// This condition takes care of wrapping.
			if nextCell == VoidCell {
				if currentDirection.IsHorizontal() {
					left, right := gridModel.GetRowEdgeCells(currentCoord[1])

					if currentDirection == DirectionRight {
						nextCoord = left
					} else {
						nextCoord = right
					}

					nextX, nextY = nextCoord[0], nextCoord[1]
					nextCell = gridModel.Grid[nextY][nextX]
				}

				if currentDirection.IsVertical() {
					top, bottom := gridModel.GetColEdgeCells(currentCoord[0])

					if currentDirection == DirectionDown {
						nextCoord = top
					} else {
						nextCoord = bottom
					}

					nextX, nextY = nextCoord[0], nextCoord[1]
					nextCell = gridModel.Grid[nextY][nextX]
				}
			}

			if nextCell == OpenCell {
				currentCoord = nextCoord
				steps = append(steps, currentCoord)
			}

			if nextCell == WallCell {
				break
			}

			//gridModel.Render(steps)
			//fmt.Print()
		}

		if i < len(instructions.Directions) {
			currentDirection = currentDirection.Rotate(instructions.Directions[i])
		}
	}

	row, col := currentCoord[1]+1, currentCoord[0]+1

	result := 1000*row + 4*col + int(currentDirection)

	fmt.Println(result)
}
