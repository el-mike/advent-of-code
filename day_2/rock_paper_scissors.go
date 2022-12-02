package day2

import "el-mike/advent-of-code/common"

const (
	INPUT_FILENAME      = "input.txt"
	TEST_INPUT_FILENAME = "test_input.txt"
)

func RockPaperScissors() {
	scanner, err := common.GetFileScanner("./day_1/" + INPUT_FILENAME)

	if err != nil {
		panic(err)
	}

	for scanner.Scan() {

	}
}
