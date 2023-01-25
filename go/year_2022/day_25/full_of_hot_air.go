package day_25

import (
	"el-mike/advent-of-code/go/common"
	"fmt"
	"strconv"
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

var SnafuDigitsNegativesMap = map[int]uint8{
	3: '=',
	4: '-',
}

func FullOfHotAir() {
	scanner, err := common.GetFileScanner("./year_2022/day_25/" + common.InputFilename)
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

	fmt.Printf("%d, snafu: %s\n", total, toSnafu(total))
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
	var snafu string

	for number != 0 {
		remainder := number % NumberBase
		number /= NumberBase

		remainderStr := strconv.Itoa(remainder)

		if remainder <= 2 {
			snafu = remainderStr + snafu
		} else {
			snafu = string(SnafuDigitsNegativesMap[remainder]) + snafu
			number += 1
		}
	}

	return snafu
}
