package common

import (
	"regexp"
	"strconv"
)

func GetNumbersFromLine(line string) ([]int, error) {
	re := regexp.MustCompile("[0-9]+")

	results := re.FindAllString(line, -1)

	var numbers []int

	for _, numberStr := range results {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}

	return numbers, nil
}
