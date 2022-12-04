package day3

import (
	"el-mike/advent-of-code/common"
)

const (
	INPUT_FILENAME      = "input.txt"
	TEST_INPUT_FILENAME = "test_input.txt"
)

const (
	ASCII_LOWERCASE_A           = 97
	LOWECASE_PRIORITY_DISTANCE  = 96
	UPPERCASE_PRIORITY_DISTANCE = 38
)

func getPriority(letter rune) int {
	asciiCode := int(letter)
	priorityDistance := 0

	// If asciiCode is equal or bigger than ASCII code for lowercase a, it's a lower case letter.
	// In order to get priority, we subtract the difference between first lowercase letter and first
	// priority value.
	if asciiCode >= ASCII_LOWERCASE_A {
		priorityDistance = LOWECASE_PRIORITY_DISTANCE
	} else {
		priorityDistance = UPPERCASE_PRIORITY_DISTANCE
	}

	return asciiCode - priorityDistance
}

func RucksackReorganization() int {
	scanner, err := common.GetFileScanner("./day_3/" + INPUT_FILENAME)
	if err != nil {
		panic(err)
	}

	totalPriority := 0

	// We use helper map to save found letters in easily readable manner.
	found := map[rune]int{}
	groupDivisor := 0

	for scanner.Scan() {
		groupDivisor += 1

		line := scanner.Text()

		// We need to find a type that appears at least once in EACH rucksack, therefore
		// we want to add the type to the found map only once per rucksack.
		checked := map[rune]bool{}

		for _, r := range line {
			if !checked[r] {
				found[r] += 1
				checked[r] = true
			}
		}

		if groupDivisor == 3 {
			for key, value := range found {
				if value == 3 {
					totalPriority += getPriority(key)
					break
				}
			}

			found = map[rune]int{}
			groupDivisor = 0
		}
	}

	return totalPriority
}
