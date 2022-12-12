package day6

import (
	"el-mike/advent-of-code/common"
	"el-mike/advent-of-code/common/ds"
)

const (
	INPUT_FILENAME      = "input.txt"
	TEST_INPUT_FILENAME = "test_input.txt"
)

const (
	MARKER_SIZE = 14
)

func TuningTrouble() int {
	scanner, err := common.GetFileScanner("./day6/" + TEST_INPUT_FILENAME)
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