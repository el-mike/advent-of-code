package day19

import "el-mike/advent-of-code/common"

func NotEnoughMinerals() {
	scanner, err := common.GetFileScanner("./day19/" + common.TestInputFilename)
	if err != nil {
		panic(err)
	}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}
	}

}
