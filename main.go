package main

import (
	"el-mike/advent-of-code/common"
	"el-mike/advent-of-code/day18"
)

func main() {
	runner := common.NewRunner()

	runner.RunAndMeasure(day18.BoilingBoulders)
}
