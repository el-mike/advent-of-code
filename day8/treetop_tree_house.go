package day8

import "el-mike/advent-of-code/common"

const (
	InputFilename     = "input.txt"
	TestInputFilename = "test_input.txt"
)

func TreetopTreeHouse() int {
	scanner, err := common.GetFileScanner("./day8/" + TestInputFilename)
	if err != nil {
		panic(err)
	}

	//var forestMap [][]int

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		//forestMap = append(forestMap, line[0])
	}

	return 0
}
