package day1

import (
	"el-mike/advent-of-code/common"
	"fmt"
	"strconv"
)

const (
	INPUT_FILENAME      = "input.txt"
	TEST_INPUT_FILENAME = "test_input.txt"
)

func Calories() {
	scanner, err := common.GetFileScanner("./day_1/" + INPUT_FILENAME)

	if err != nil {
		panic(err)
	}

	result := []int{}
	current := 0

	for scanner.Scan() {

		line := scanner.Text()

		if line != "" {
			value, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}

			current += value
		} else {
			result = append(result, current)
			current = 0
		}
	}

	// If the file does not end with a line with only newline character,
	// the last amount would not be saved.
	if current != 0 {
		result = append(result, current)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	common.QuickSort(result, 0, (len(result) - 1))

	bestThree := result[len(result)-3:]

	sum := 0

	for _, x := range bestThree {
		sum += x
	}

	fmt.Println(sum)
}
