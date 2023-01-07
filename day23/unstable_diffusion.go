package day23

import (
	"el-mike/advent-of-code/common"
	"el-mike/advent-of-code/common/ds"
	"fmt"
)

func UnstableDiffustion() {
	scanner, err := common.GetFileScanner("./day23/" + common.InputFilename)
	if err != nil {
		panic(err)
	}

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		lines = append(lines, line)
	}

	gridModel := NewGridModel()
	gridModel.Build(lines)

	directionsQueue := ds.NewQueue[Direction]()

	directionsQueue.Enqueue(DirectionN)
	directionsQueue.Enqueue(DirectionS)
	directionsQueue.Enqueue(DirectionW)
	directionsQueue.Enqueue(DirectionE)

	round := 1

	for ; ; round++ {
		propositionsMap := map[string][]Coord{}
		directions := getDirectionsAndRequeue(directionsQueue)

		for _, position := range gridModel.ElfPositions {
			if !gridModel.HasAnyNeighbor(position) {
				continue
			}

			var proposition Coord

			for _, direction := range directions {
				hasNeighbor, err := gridModel.HasNeighborByDirection(position, direction)
				if err != nil {
					fmt.Println(err)
				}

				if !hasNeighbor {
					proposition = position.GetByDirection(direction)
					break
				}
			}

			if proposition != NilCoord() {
				key := proposition.GetKey()
				propositionsMap[key] = append(propositionsMap[key], position)
			}
		}

		if len(propositionsMap) == 0 {
			break
		}

		padTop, padRight, padBottom, padLeft := false, false, false, false

		for key, propositions := range propositionsMap {
			if len(propositions) == 1 {
				currentCoord := propositions[0]
				nextCoord := NewCoordFromKey(key)

				delete(gridModel.ElfPositions, currentCoord.GetKey())

				gridModel.ElfPositions[nextCoord.GetKey()] = nextCoord

				padTop = padTop || gridModel.ExtendsTop(nextCoord)
				padRight = padRight || gridModel.ExtendsRight(nextCoord)
				padBottom = padBottom || gridModel.ExtendsBottom(nextCoord)
				padLeft = padLeft || gridModel.ExtendsLeft(nextCoord)
			}
		}

		gridModel.PadIfExtends(padTop, padRight, padBottom, padLeft)
	}

	fmt.Println(round)
}

func getDirectionsAndRequeue(queue *ds.Queue[Direction]) []Direction {
	dir1, err := queue.Dequeue()
	if err != nil {
		panic(err)
	}

	dir2, err := queue.Dequeue()
	if err != nil {
		panic(err)
	}

	dir3, err := queue.Dequeue()
	if err != nil {
		panic(err)
	}

	dir4, err := queue.Dequeue()
	if err != nil {
		panic(err)
	}

	directions := []Direction{dir1, dir2, dir3, dir4}

	queue.Enqueue(dir2)
	queue.Enqueue(dir3)
	queue.Enqueue(dir4)
	queue.Enqueue(dir1)

	return directions
}
