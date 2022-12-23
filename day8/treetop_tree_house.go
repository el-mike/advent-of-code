package day8

import (
	"el-mike/advent-of-code/common"
	"fmt"
)

func TreetopTreeHouse() int {
	scanner, err := common.GetFileScanner("./day8/" + common.InputFilename)
	if err != nil {
		panic(err)
	}

	var forestMap [][]int
	i := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		forestMap = append(forestMap, []int{})

		for _, r := range line {
			// Trivial way to get int value from rune containing ASCII representation
			forestMap[i] = append(forestMap[i], int(r-'0'))
		}

		i += 1
	}

	numRows := len(forestMap)
	numCols := len(forestMap[0])

	max := 0

	for i = 0; i < numRows; i += 1 {
		for j := 0; j < numCols; j += 1 {
			current := forestMap[i][j]

			numLeft, numRight, numUp, numDown := 0, 0, 0, 0

			for k := i - 1; k >= 0; k -= 1 {
				numUp += 1

				if forestMap[k][j] >= current {
					break
				}
			}

			for k := i + 1; k < numRows; k += 1 {
				numDown += 1

				if forestMap[k][j] >= current {
					break
				}
			}

			for k := j - 1; k >= 0; k -= 1 {
				numLeft += 1

				if forestMap[i][k] >= current {
					break
				}
			}

			for k := j + 1; k < numCols; k += 1 {
				numRight += 1

				if forestMap[i][k] >= current {
					break
				}
			}

			score := numLeft * numRight * numUp * numDown

			if score > max {
				max = score
			}
		}
	}
	return max
}

func renderForest(forestMap [][]int, markY, markX int) {
	for i := 0; i < len(forestMap); i += 1 {
		for j := 0; j < len(forestMap[0]); j += 1 {
			if i == markY && j == markX {
				fmt.Print("\033[31m")
			}

			fmt.Printf("%d", forestMap[i][j])
			fmt.Print("\033[0m")
		}

		fmt.Printf("\n")
	}

	fmt.Printf("\n")
}
