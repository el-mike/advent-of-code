package day_25

import (
	"el-mike/advent-of-code/go/common"
	"fmt"
)

const (
	NumberBase = 5
)

var SnafuDigitsMap = map[uint8]int{
	'=': -2,
	'-': -1,
	'0': 0,
	'1': 1,
	'2': 2,
}

func FullOfHotAir() {
	scanner, err := common.GetFileScanner("./year_2022/day_25/" + common.TestInputFilename)
	if err != nil {
		panic(err)
	}

	var snafuNumbers []string

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		snafuNumbers = append(snafuNumbers, line)
	}

	total := 0

	for _, snafu := range snafuNumbers {
		total += fromSnafu(snafu)
	}

	fmt.Println(total)

	fmt.Println(toSnafu(total))
}

func fromSnafu(snafu string) int {
	number := 0
	coefficient := 1

	for i := (len(snafu) - 1); i >= 0; i-- {
		r := snafu[i]

		number += SnafuDigitsMap[r] * coefficient

		coefficient *= NumberBase
	}

	return number
}

func toSnafu(number int) string {
	snafu := ""

	for ; number != 0; {
		snafu = fmt.Sprintf("%d%s", (number % NumberBase), snafu)
		number /= NumberBase
	}

	return snafu
}
