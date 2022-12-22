package main

import (
	"el-mike/advent-of-code/common"
	"el-mike/advent-of-code/day17"
)

func main() {
	runner := common.NewRunner()

	runner.RunAndMeasure(day17.PyroclasticFlow)
}
