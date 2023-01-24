package day_24

import (
	"el-mike/advent-of-code/go/common"
	"el-mike/advent-of-code/go/common/ds"
	"fmt"
)

type StepInfo struct {
	Minute   int
	Position Vector
}

func BlizzardBasin() {
	scanner, err := common.GetFileScanner("./year_2022/day_24/" + common.InputFilename)
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

	period := common.LCM(gridModel.Width-2, gridModel.Height-2)

	blizzardStates := map[int]BlizzardPositions{}

	for i := 0; i < period; i++ {
		blizzardStates[i] = gridModel.BlizzardPositions
		gridModel.MoveBlizzards()
	}

	frontier := ds.NewQueue[*StepInfo]()
	visited := map[string]*StepInfo{}

	frontier.Enqueue(&StepInfo{Minute: 0, Position: gridModel.Start})

	var lastStep *StepInfo

Outer:
	for !frontier.IsEmpty() {
		current, err := frontier.Dequeue()
		if err != nil {
			panic(err)
		}

		visited[current.Position.ID()] = current

		nextMinute := current.Minute + 1
		gridModel.BlizzardPositions = blizzardStates[nextMinute%period]

		// We add current.Position as a "wait" step.
		candidates := append(current.Position.GetNeighbors(), current.Position)

		for _, candidate := range candidates {
			candidateStep := &StepInfo{Minute: nextMinute, Position: candidate}

			if gridModel.IsEnd(candidate) {
				lastStep = candidateStep
				break Outer
			}

			if !gridModel.IsInBounds(candidate) ||
				gridModel.IsWall(candidate) ||
				gridModel.hasBlizzard(candidate) {
				continue
			}

			frontier.Enqueue(candidateStep)
		}

	}

	fmt.Println(lastStep.Minute)
}
