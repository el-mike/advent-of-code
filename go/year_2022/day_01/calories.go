package day_01

import (
	"el-mike/advent-of-code/go/common"
	"fmt"
	"strconv"
)

func Calories() int {
	scanner, err := common.GetFileScanner("./year_2022/day_01/" + common.InputFilename)

	if err != nil {
		panic(err)
	}

	var result []int
	current := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			value, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}

			current += value
		} else {
			result = append(result, current)
			current = 0
		}
	}

	// If the file does not end with a line with only newline character,
	// the last amount would not be saved.
	if current != 0 {
		result = append(result, current)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	common.QuickSort(result, 0, (len(result) - 1))

	bestThree := result[len(result)-3:]

	sum := 0

	for _, x := range bestThree {
		sum += x
	}

	return sum
}

func SonarSweep() {
	scanner, err := common.GetFileScanner("../rust/src/year_2021/day_01/" + common.InputFilename)
	if err != nil {
		panic(err)
	}

	count := 0
	previous := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			value, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}

			if previous == 0 {
				previous = value
				continue
			}

			if value > previous {
				count += 1
			}

			previous = value
		}
	}

	fmt.Println(count)
}
