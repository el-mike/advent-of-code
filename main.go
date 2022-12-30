package main

import (
	"el-mike/advent-of-code/common"
	"el-mike/advent-of-code/day19"
)

func main() {
	runner := common.NewRunner()

	runner.RunAndMeasure(day19.NotEnoughMinerals)
}
