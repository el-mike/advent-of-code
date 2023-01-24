package main

import (
	"el-mike/advent-of-code/go/common"
	"el-mike/advent-of-code/go/year_2022/day_25"
)

func main() {
	runner := common.NewRunner()

	runner.RunAndMeasure(day_25.FullOfHotAir)
}
