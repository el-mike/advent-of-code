package day6

import (
	"el-mike/advent-of-code/common"
)

const (
	INPUT_FILENAME      = "input.txt"
	TEST_INPUT_FILENAME = "test_input.txt"
)

const (
	MARKER_SIZE = 14
)

func TuningTrouble() int {
	scanner, err := common.GetFileScanner("./day6/" + INPUT_FILENAME)
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

	markerBuffer := make([]rune, MARKER_SIZE)
	charsMap := map[rune]int{}

	for i, char := range input {
		lastIndex := i % MARKER_SIZE

		if i >= MARKER_SIZE {
			charsMap[markerBuffer[lastIndex]] -= 1
		}

		markerBuffer[lastIndex] = char
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
