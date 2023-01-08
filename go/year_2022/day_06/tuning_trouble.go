package day_06

import (
	"el-mike/advent-of-code/go/common"
	"el-mike/advent-of-code/go/common/ds"
)

const (
	MARKER_SIZE = 14
)

func TuningTrouble() int {
	scanner, err := common.GetFileScanner("./year_2022/day_06/" + common.InputFilename)
	if err != nil {
		panic(err)
	}

	input := ""

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		input += line
	}

	charsQueue := ds.NewQueue[rune]()
	charsMap := map[rune]int{}

	for i, char := range input {
		if i >= MARKER_SIZE {
			lastChar, err := charsQueue.Dequeue()
			if err != nil {
				panic(err)
			}

			charsMap[lastChar] -= 1
		}

		charsQueue.Enqueue(char)
		charsMap[char] += 1

		if i >= MARKER_SIZE &&
			!common.AnySatisfies(
				charsMap,
				func(_ rune, value int) bool {
					return value >= 2
				},
			) {
			// We need to add 1 to the index, as we want to return the number of the position
			// after which message package starts.
			return (i + 1)
		}

	}

	return 0
}
