package day21

import (
	"el-mike/advent-of-code/common"
	"fmt"
	"math"
	"regexp"
	"strings"
)

type MonkeysMap map[string]*Monkey

const RootMonkeyName = "root"

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
	scanner, err := common.GetFileScanner("./day21/" + common.InputFilename)
	if err != nil {
		panic(err)
	}

	monkeysMap := MonkeysMap{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		monkey := parseLine(line)

		monkeysMap[monkey.Name] = monkey
	}

	result := getNumber(monkeysMap, monkeysMap[RootMonkeyName])

	fmt.Println(result)
}

func getNumber(
	monkeysMap MonkeysMap,
	monkey *Monkey,
) int {
	if monkey.HasNumber() {
		return monkey.Number
	}

	leftMonkey, rightMonkey := monkeysMap[monkey.LeftMonkey], monkeysMap[monkey.RightMonkey]

	if !leftMonkey.HasNumber() {
		leftMonkey.Number = getNumber(monkeysMap, leftMonkey)
	}

	if !rightMonkey.HasNumber() {
		rightMonkey.Number = getNumber(monkeysMap, rightMonkey)
	}

	monkey.Number = monkey.OperationFn(leftMonkey.Number, rightMonkey.Number)

	return monkey.Number
}
