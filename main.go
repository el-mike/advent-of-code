package main

import (
	"el-mike/advent-of-code/common"
	"el-mike/advent-of-code/day23"
)

func main() {
	runner := common.NewRunner()

	runner.RunAndMeasure(day23.UnstableDiffustion)
}
