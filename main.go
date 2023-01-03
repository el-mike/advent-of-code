package main

import (
	"el-mike/advent-of-code/common"
	"el-mike/advent-of-code/day20"
)

func main() {
	runner := common.NewRunner()

	runner.RunAndMeasure(day20.GrovePositioningSystem)
}
