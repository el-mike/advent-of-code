package day20

import (
	"el-mike/advent-of-code/common"
	"fmt"
)

func GrovePositioningSystem() {
	scanner, err := common.GetFileScanner("./day20/" + common.InputFilename)
	if err != nil {
		panic(err)
	}

	var origin []int
	var mixed []int
	positionMap := map[int]int{}

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		if line == "" {
			continue
		}

		numbers, err := common.GetNumbersFromLine(line)
		if err != nil {
			panic(err)
		}

		number := numbers[0]

		origin = append(origin, number)
		mixed = append(mixed, number)

		// At the beginning, every index corresponds to itself.
		positionMap[i] = i
	}

	length := len(origin)

	for i, value := range origin {
		if value == 0 {
			continue
		}

		source := positionMap[i]

		var dest int

		nextPosition := source + value

		if value < 0 {
			nextPosition -= 1
		}

		if nextPosition >= 0 {
			dest = nextPosition % (length - 1)
		} else {
			// We need to use addition instead of subtraction, as result of
			// module will be negative as well.
			dest = length + (nextPosition % (length - 1))
		}

		if dest == length {
			dest = 0
		}

		common.Move[int](mixed, source, dest)

		var start, end int

		if dest > source {
			start, end = (source + 1), dest
		} else {
			start, end = dest, (source - 1)
		}

		indexesToUpdate := getIndexesToUpdate(positionMap, start, end)

		for _, index := range indexesToUpdate {
			if dest > source {
				positionMap[index] -= 1
			} else {
				positionMap[index] += 1
			}
		}

		positionMap[i] = dest
	}

	var zeroIndex int

	for i, value := range mixed {
		if value == 0 {
			zeroIndex = i
			break
		}
	}

	coordinates := []int{
		mixed[(zeroIndex+1000)%length],
		mixed[(zeroIndex+2000)%length],
		mixed[(zeroIndex+3000)%length],
	}

	result := coordinates[0] + coordinates[1] + coordinates[2]

	fmt.Println(result)
}

func getIndexesToUpdate(m map[int]int, start, end int) []int {
	var indexes []int

	for key, value := range m {
		if value >= start && value <= end {
			indexes = append(indexes, key)
		}
	}

	return indexes
}
