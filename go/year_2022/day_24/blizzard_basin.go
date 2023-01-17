package day_24

import (
	"el-mike/advent-of-code/go/common"
	"fmt"
)

func BlizzardBasin() {
	scanner, err := common.GetFileScanner("./year_2022/day_24/" + common.TestInputFilename)
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

	currentPosition := gridModel.Start

	for {
		gridModel.Render(currentPosition)
		gridModel.MoveBlizzards()
	}

	fmt.Println()
}
