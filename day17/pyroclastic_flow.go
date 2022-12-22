package day17

import (
	"el-mike/advent-of-code/common"
)

const (
	InputFilename     = "input.txt"
	TestInputFilename = "test_input.txt"
)

func PyroclasticFlow() {
	scanner, err := common.GetFileScanner("./day/17" + TestInputFilename)
	if err != nil {
		panic(err)
	}

	for scanner.Scan() {

	}
}
