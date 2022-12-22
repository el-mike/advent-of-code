package day17

import (
	"el-mike/advent-of-code/common"
	"fmt"
)

const (
	InputFilename     = "input.txt"
	TestInputFilename = "test_input.txt"
)

const (
	NumRockTypes = 5
	// A trillion.
	//NumRocks = 1000000000000
	NumRocks = 2022
)

var rockTypesMap = map[int]RockType{
	0: Horizontal,
	1: Cross,
	2: InvertedL,
	3: Vertical,
	4: Square,
}

func PyroclasticFlow() {
	scanner, err := common.GetFileScanner("./day17/" + InputFilename)
	if err != nil {
		panic(err)
	}

	directions := Directions{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		directions.FromString(line)
	}

	board := NewBoard()

	directionIndex := 0

	for i := 0; i < NumRocks; i += 1 {
		currentRock := NewRock(board, rockTypesMap[i%NumRockTypes])

		board.AddRock(currentRock)

		for j := 0; ; j += 1 {
			direction := directions[directionIndex%len(directions)]

			// First, we move sideways.
			currentRock.Move(direction)

			//Then we move the rock down.
			settled := currentRock.Move(Down)

			directionIndex += 1
			if settled {
				//tryRender(board)
				break
			}
		}
	}

	fmt.Println(board.ReachedY + 1)

	return
}

func tryRender(board *Board) {
	common.ClearTerminal()
	board.Render()
	// time.Sleep(500 * time.Millisecond)
}
