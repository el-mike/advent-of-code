package day_21

import (
	"el-mike/advent-of-code/go/common"
	"fmt"
	"math"
	"regexp"
	"strings"
)

type MonkeysMap map[string]*Monkey

const (
	RootMonkeyName = "root"
	MyMonkeyName   = "humn"
)

// Arbitrary values to start from. Its size is based on empirical testing.
const StartTestValue = 10000
const StartFactor = 12

func parseLine(line string) *Monkey {
	parts := strings.Split(line, ":")

	name := strings.TrimSpace(parts[0])
	job := strings.TrimSpace(parts[1])

	var number int = math.MaxInt
	var operation, leftMonkey, rightMonkey string

	re := regexp.MustCompile(`(.*)\s([+\-*/])\s(.*)`)

	if re.MatchString(job) {
		jobParts := re.FindAllStringSubmatch(job, -1)[0]

		leftMonkey = jobParts[1]
		rightMonkey = jobParts[3]
		operation = jobParts[2]
	} else {
		numbers, err := common.GetNumbersFromLine(line)
		if err != nil {
			panic(err)
		}

		number = numbers[0]
	}

	return NewMonkey(name, number, leftMonkey, rightMonkey, operation)
}

func MonkeyMath() {
	scanner, err := common.GetFileScanner("./year_2022/day_21/" + common.InputFilename)
	if err != nil {
		panic(err)
	}

	originalMonkeysMap := MonkeysMap{}
	currentMonkeysMap := MonkeysMap{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		monkey := parseLine(line)

		originalMonkeysMap[monkey.Name] = monkey
		currentMonkeysMap[monkey.Name] = monkey.Clone()
	}

	rootMonkey := originalMonkeysMap[RootMonkeyName]
	testValue := StartTestValue

	var left, right, lastGreater int
	currentFactor := StartFactor

	for {
		// @TODO:
		// Unfortunately, because of some precision error, there are more than one
		// testValue that makes the test pass. We can find the lowest one by picking the found
		// value and changing it to find the correct answer, but the precision error should be fixed,
		// to avoid this manual work.
		testValue += int(math.Pow(10, float64(currentFactor)))
		leftMonkey, rightMonkey := currentMonkeysMap[rootMonkey.LeftMonkey], currentMonkeysMap[rootMonkey.RightMonkey]

		left = getNumber(currentMonkeysMap, leftMonkey, testValue)
		right = getNumber(currentMonkeysMap, rightMonkey, testValue)

		if left == right {
			break
		}

		diff := left - right

		// Based on testing, the left value is converging to right value
		// when testValue goes up.
		if diff > 0 {
			lastGreater = testValue
		} else {
			if currentFactor > 0 {
				currentFactor -= 1
			}

			if lastGreater > 0 {
				testValue = lastGreater
			}
		}

		resetMap(originalMonkeysMap, currentMonkeysMap)
	}

	fmt.Println(testValue)
}

func getNumber(
	monkeysMap MonkeysMap,
	monkey *Monkey,
	testValue int,
) int {
	if monkey.Name == MyMonkeyName {
		return testValue
	}
	if monkey.HasNumber() {
		return monkey.Number
	}

	left := getNumber(monkeysMap, monkeysMap[monkey.LeftMonkey], testValue)
	right := getNumber(monkeysMap, monkeysMap[monkey.RightMonkey], testValue)

	monkey.Number = monkey.OperationFn(left, right)

	return monkey.Number
}

func resetMap(original, current MonkeysMap) {
	for key, monkey := range original {
		current[key].Number = monkey.Number
	}
}
