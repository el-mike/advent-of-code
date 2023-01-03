package main

import (
	"el-mike/advent-of-code/common"
	"el-mike/advent-of-code/day21"
)

func main() {
	runner := common.NewRunner()

	runner.RunAndMeasure(day21.MonkeyMath)
}
