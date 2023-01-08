package day_09

import (
	"el-mike/advent-of-code/go/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	Up = iota
	Down
	Left
	Right
)

const (
	UpSign    = "U"
	DownSign  = "D"
	LeftSign  = "L"
	RightSign = "R"
)

var differenceMap = map[int]int{
	Up:    -1,
	Down:  1,
	Left:  -1,
	Right: 1,
}

var directionsMap = map[string]int{
	UpSign:    Up,
	DownSign:  Down,
	LeftSign:  Left,
	RightSign: Right,
}

const BodyLength = 10

func parseLine(line string) (int, int) {
	parts := strings.Split(line, " ")

	direction := directionsMap[parts[0]]
	numSteps, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return direction, numSteps
}

func isAdjacent(leadPos, followPos []int) bool {
	return math.Abs(float64(leadPos[0])-float64(followPos[0])) <= 1 &&
		math.Abs(float64(leadPos[1])-float64(followPos[1])) <= 1
}

func handleHead(head []int, direction int) {
	diff := differenceMap[direction]

	if direction == Up || direction == Down {
		head[0] += diff
	} else {
		head[1] += diff
	}
}

func handleBody(leadPos, followPos []int) {
	if !isAdjacent(leadPos, followPos) {
		if leadPos[0] > followPos[0] {
			followPos[0] += differenceMap[Down]
		}
		if leadPos[0] < followPos[0] {
			followPos[0] += differenceMap[Up]
		}
		if leadPos[1] > followPos[1] {
			followPos[1] += differenceMap[Right]
		}
		if leadPos[1] < followPos[1] {
			followPos[1] += differenceMap[Left]
		}
	}
}

func getKey(tailPos []int) string {
	return fmt.Sprintf("[%d][%d]", tailPos[0], tailPos[1])
}

func RopeBridge() int {
	scanner, err := common.GetFileScanner("./year_2022/day_09/" + common.InputFilename)
	if err != nil {
		panic(err)
	}

	renderer := NewBoardRenderer()

	visitedMap := map[string]bool{}

	bodyPos := make([][]int, BodyLength)

	for i := 0; i < BodyLength; i += 1 {
		bodyPos[i] = []int{0, 0}
	}

	visitedMap[getKey(bodyPos[BodyLength-1])] = true

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		direction, numSteps := parseLine(line)

		for i := 0; i < numSteps; i += 1 {
			for j := 0; j < (BodyLength - 1); j += 1 {
				if j == 0 {
					handleHead(bodyPos[j], direction)
				}

				handleBody(bodyPos[j], bodyPos[j+1])
			}

			visitedMap[getKey(bodyPos[BodyLength-1])] = true
		}

		renderer.render(bodyPos)
	}

	return common.CountWhere(visitedMap, func(key string, value bool) bool {
		return value
	})
}
