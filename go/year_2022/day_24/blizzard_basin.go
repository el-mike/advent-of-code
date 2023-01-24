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

func (si *StepInfo) ID() string {
	return fmt.Sprintf("%d|%s", si.Minute, si.Position.ID())
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

	for !frontier.IsEmpty() {
		current, err := frontier.Dequeue()
		if err != nil {
			panic(err)
		}

		if gridModel.IsEnd(current.Position) {
			lastStep = current
			break
		}

		// Checking visited cells is necessary, as otherwise search space grows
		// very fast. However, we need to include Minute in the ID function as well,
		// otherwise it will allow to stay in place only one time.
		if _, ok := visited[current.ID()]; ok {
			continue
		}

		visited[current.ID()] = current

		nextMinute := current.Minute + 1
		gridModel.BlizzardPositions = blizzardStates[nextMinute%period]

		// We add current.Position as a "wait" step.
		candidates := append(current.Position.GetNeighbors(), current.Position)

		for _, candidate := range candidates {
			if !gridModel.IsInBounds(candidate) ||
				gridModel.IsWall(candidate) ||
				gridModel.hasBlizzard(candidate) {
				continue
			}

			frontier.Enqueue(&StepInfo{Minute: nextMinute, Position: candidate})
		}

	}

	fmt.Println(lastStep.Minute)
}
