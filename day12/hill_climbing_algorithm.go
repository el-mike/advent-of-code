package day12

import (
	"el-mike/advent-of-code/common"
	"el-mike/advent-of-code/common/ds"
	"fmt"
)

const (
	InputFilename     = "input.txt"
	TestInputFilename = "test_input.txt"
)

func HillClimbingAlgorithm() {
	scanner, err := common.GetFileScanner("./day12/" + InputFilename)
	if err != nil {
		panic(err)
	}

	var gridStr []string

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		gridStr = append(gridStr, line)
	}

	grid := NewGridModel(len(gridStr), len(gridStr[0]))
	grid.Parse(gridStr)

	var possibleStarts []*GridCell

	grid.ForEach(func(cell *GridCell, _, _ int) {
		if cell.Char == 'S' || cell.Char == 'a' {
			possibleStarts = append(possibleStarts, cell)
		}
	})

	var distances []int

	for _, start := range possibleStarts {
		path := ds.NewStack[*GridCell]()

		stepsMap := BreadthFirstSearch(grid, start)

		current := grid.End

		for current != nil && !current.Same(start) {
			path.Push(current)
			current = stepsMap[current.ID()]
		}

		// If current equals nil, it means that the path does not exist.
		if current == nil {
			continue
		}

		distances = append(distances, path.Size())
	}

	min := distances[0]
	for _, distance := range distances {
		if distance < min {
			min = distance
		}
	}

	fmt.Print(min)
}
