package day22

import (
	"el-mike/advent-of-code/common"
	"fmt"
)

const FaceSize = 50

//const FaceSize = 4

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

	gridModel := NewGridModel(FaceSize)
	instructions := NewInstructions(movesStr)

	gridModel.Build(lines, maxRow)

	currentCoord := Coord{0, 0}

	for x, cell := range gridModel.Grid[0] {
		if cell != VoidCell {
			currentCoord[0] = x
			break
		}
	}

	//currentCoord = Coord{currentCoord[0] + 1, 0}

	steps := []Coord{currentCoord}

	currentDirection := DirectionRight
	//currentDirection := DirectionUp

	for i := 0; i < len(instructions.Numbers); i++ {
		n := instructions.Numbers[i]

		for j := 0; j < n; j++ {
			var nextCoord Coord
			// We could use any value bigger than 3, as there are only
			// 4 directions indexed from 0 to 3, so 4 is the first non-existing direction.
			nextDirection := Direction(4)

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
				nextCoord, nextDirection = gridModel.GetWrapped(currentCoord, currentDirection)

				nextX, nextY = nextCoord[0], nextCoord[1]
				nextCell = gridModel.Grid[nextY][nextX]
			}

			if nextCell == OpenCell {
				currentCoord = nextCoord
				steps = append(steps, currentCoord)
			}

			if nextCell == WallCell {
				break
			}

			// If new nextDirection was returned by GetWrapper and
			// next cell is not a WallCell, replace current direction.
			if nextDirection != 4 {
				currentDirection = nextDirection
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
