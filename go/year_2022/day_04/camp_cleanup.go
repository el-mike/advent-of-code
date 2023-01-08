package day_04

import (
	"el-mike/advent-of-code/go/common"
	"strings"
)

func CampCleanup() int {
	scanner, err := common.GetFileScanner("./year_2022/day_04/" + common.InputFilename)
	if err != nil {
		panic(err)
	}

	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		assignmentPairs := strings.Split(line, ",")

		firstAssignment := NewAssignment(assignmentPairs[0])
		secondAssignment := NewAssignment(assignmentPairs[1])

		if firstAssignment.Overlaps(secondAssignment) {
			count += 1
		}
	}

	return count
}
