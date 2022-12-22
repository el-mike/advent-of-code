package main

import (
	"el-mike/advent-of-code/common"
	"el-mike/advent-of-code/day16"
)

func main() {
	runner := common.NewRunner()

	runner.RunAndMeasure(day16.ProboscideaVolcanium)
}
