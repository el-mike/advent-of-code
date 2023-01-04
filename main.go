package main

import (
	"el-mike/advent-of-code/common"
	"el-mike/advent-of-code/day22"
)

func main() {
	runner := common.NewRunner()

	runner.RunAndMeasure(day22.MonkeyMap)
}
