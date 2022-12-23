package day12

import (
	"el-mike/advent-of-code/common"
	"el-mike/advent-of-code/common/ds"
	"fmt"
)

func HillClimbingAlgorithm() {
	scanner, err := common.GetFileScanner("./day12/" + common.InputFilename)
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

	grid := NewGridModel(gridStr)

	var possibleStarts []*GridCell

	grid.ForEach(func(cell *GridCell, _, _ int) {
		if cell.Char == 'S' || cell.Char == 'a' {
			possibleStarts = append(possibleStarts, cell)
		}
	})

	var paths []*ds.Stack[*GridCell]

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

		paths = append(paths, path)
	}

	min := paths[0]
	for _, path := range paths {
		if path.Size() < min.Size() {
			min = path
		}
	}

	fmt.Print(min.Size())

	var pathSlice []*GridCell

	for !min.Empty() {
		pathCell, err := min.Pop()
		if err != nil {
			panic(err)
		}

		pathSlice = append(pathSlice, pathCell)
	}

	grid.Render(pathSlice)
}
