package day17

import (
	"el-mike/advent-of-code/common"
	"fmt"
	"time"
)

type IterationSnapshot struct {
	LastDirectionIndex int
	LastRockTypeIndex  int
	SettledRocks       int
	ReachedY           int
}

type StatesMap map[int]map[string]*IterationSnapshot

type CycleResult struct {
	PreCycleSnapshot *IterationSnapshot
	CycleSnapshot    *IterationSnapshot
}

const (
	NumRockTypes = 5
	// A trillion.
	NumRocks        = 1000000000000
	NumMemoizedRows = 100
)

var rockTypesMap = map[int]RockType{
	0: Horizontal,
	1: Cross,
	2: InvertedL,
	3: Vertical,
	4: Square,
}

func PyroclasticFlow() {
	scanner, err := common.GetFileScanner("./day17/" + common.TestInputFilename)
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

	board := NewBoard(nil)

	cycleResult := singleRun(board, directions, true, NumRocks, 0, 0)

	preCycleSnapshot := cycleResult.PreCycleSnapshot
	cycleSnapshot := cycleResult.CycleSnapshot

	reachedYPreCycle := preCycleSnapshot.ReachedY

	rocksLeftForCycles := NumRocks - preCycleSnapshot.SettledRocks
	rocksSettledInCycle := cycleSnapshot.SettledRocks - preCycleSnapshot.SettledRocks
	reachedYForCycle := cycleSnapshot.ReachedY - preCycleSnapshot.ReachedY

	// How many full cycles will fit in the remaining rocks to simulate.
	numCycles := rocksLeftForCycles / rocksSettledInCycle

	// How many rocks are left to check (the last round).
	rocksLeft := rocksLeftForCycles % rocksSettledInCycle

	// ReachedY from cycles.
	cyclesReachedY := numCycles * reachedYForCycle

	// Top row from the initial board (when we found cycle and ended the run)
	// should be the floor for the final round.

	board = NewBoard(board.GetTopRows(NumMemoizedRows))

	// Since we ended of those values in the cycle, we need to start with the
	// consecutive values.
	startDirectionIndex := cycleSnapshot.LastDirectionIndex + 1
	startRockTypeIndex := cycleSnapshot.LastRockTypeIndex + 1

	singleRun(board, directions, false, rocksLeft, startDirectionIndex, startRockTypeIndex)

	lastRoundReachedY := board.ReachedY - board.InitialReachY

	// We need to add one, as reachedY is always an index, and the result should be the
	// height (length).
	totalHeight := reachedYPreCycle + cyclesReachedY + lastRoundReachedY + 1

	fmt.Println(totalHeight)

	return
}

func singleRun(
	board *Board,
	directions Directions,
	firstCycleOnly bool,
	numRocks int,
	directionIndex int,
	rockTypeIndex int,
) *CycleResult {
	settledRocks := 0

	round := 0
	statesMap := StatesMap{}

	for i := 0; i < numRocks; i += 1 {
		currentRock := NewRock(board, rockTypesMap[rockTypeIndex])

		board.AddRock(currentRock)

		for {
			direction := directions[directionIndex]

			// First, we move sideways.
			currentRock.Move(direction)

			//Then we move the rock down.
			settled := currentRock.Move(Down)

			directionIndex = (directionIndex + 1) % len(directions)

			if directionIndex == 0 {
				round += 1
			}

			if settled {
				settledRocks += 1

				if firstCycleOnly {
					key := getPatternKey(rockTypeIndex, directionIndex, board)

					for roundIndex := range statesMap {
						for stateKey := range statesMap[roundIndex] {
							if stateKey == key {
								preCycleStatus := statesMap[roundIndex][stateKey]

								cycleStatus := &IterationSnapshot{
									LastDirectionIndex: directionIndex,
									LastRockTypeIndex:  rockTypeIndex,
									SettledRocks:       settledRocks,
									ReachedY:           board.ReachedY,
								}

								return &CycleResult{
									PreCycleSnapshot: preCycleStatus,
									CycleSnapshot:    cycleStatus,
								}
							}
						}
					}

					if _, ok := statesMap[round]; !ok {
						statesMap[round] = map[string]*IterationSnapshot{}
					}

					statesMap[round][key] = &IterationSnapshot{
						LastDirectionIndex: directionIndex,
						LastRockTypeIndex:  rockTypeIndex,
						SettledRocks:       settledRocks,
						ReachedY:           board.ReachedY,
					}
				}

				break
			}
		}

		rockTypeIndex = (rockTypeIndex + 1) % NumRockTypes
	}

	return nil
}

func getPatternKey(
	rockTypeIndex int,
	directionIndex int,
	board *Board,
) string {
	rowStr := ""

	for _, row := range board.GetTopRows(NumMemoizedRows) {
		for _, cell := range row {
			rowStr += cellSignsMap[cell]
		}
	}

	return fmt.Sprintf("%d|%d|%s", rockTypeIndex, directionIndex, rowStr)
}

func tryRender(board *Board) {
	common.ClearTerminal()
	board.Render()
	time.Sleep(500 * time.Millisecond)
}
