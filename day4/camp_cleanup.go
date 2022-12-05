package day4

import (
	"el-mike/advent-of-code/common"
	"strings"
)

const (
	INPUT_FILENAME      = "input.txt"
	TEST_INPUT_FILENAME = "test_input.txt"
)

func CampCleanup() int {
	scanner, err := common.GetFileScanner("./day4/" + INPUT_FILENAME)
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
