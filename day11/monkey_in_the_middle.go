package day11

import (
	"el-mike/advent-of-code/common"
	"fmt"
	"regexp"
	"strings"
)

const (
	NumRounds = 10000
)

func parseOperation(line string) []string {
	equation := strings.Split(line, "=")[1]
	equation = strings.TrimSpace(equation)

	// This regex captures:
	// 	1. Any non-white-space character
	//  2. Operator
	//  3. Operand
	// \s need to be added between operator and left/right, as otherwise the matches won't work.
	re := regexp.MustCompile(`(\S*)\s(\+|-|\*|/)\s(.*)`)

	results := re.FindAllStringSubmatch(equation, -1)[0]

	// index "0" contains entire match.
	return []string{results[1], results[2], results[3]}
}

func MonkeyInTheMiddle() {
	scanner, err := common.GetFileScanner("./day11/" + common.InputFilename)
	if err != nil {
		panic(err)
	}

	monkeysMap := map[int]*Monkey{}
	var monkeys []*Monkey

	monkey := NewMonkey()

	for i := 0; scanner.Scan(); i += 1 {
		line := scanner.Text()

		// Empty lines divides each monkey's input data.
		if line == "" {
			monkeysMap[monkey.ID] = monkey
			monkeys = append(monkeys, monkey)
			monkey = NewMonkey()

			continue
		}

		mod := i % 7

		numbers, err := common.GetNumbersFromLine(line)
		if err != nil {
			panic(err)
		}

		switch mod {
		case 0:
			monkey.ID = numbers[0]
		case 1:
			for _, number := range numbers {
				monkey.Items = append(monkey.Items, number)
			}
		case 2:
			parts := parseOperation(line)

			_, operator, right := parts[0], parts[1], parts[2]

			monkey.Operation = OperationsMap[operator]
			monkey.Operand = right
		case 3:
			monkey.TestValue = numbers[0]
		case 4:
			fallthrough
		case 5:
			monkey.Recipients = append(monkey.Recipients, numbers[0])
		}
	}

	monkeysMap[monkey.ID] = monkey
	monkeys = append(monkeys, monkey)

	max := 1

	for _, monkey = range monkeys {
		max *= monkey.TestValue
	}

	for i := 0; i < NumRounds; i += 1 {
		for _, monkey = range monkeys {
			if len(monkey.Items) == 0 {
				continue
			}

			trueId, falseId := monkey.Recipients[0], monkey.Recipients[1]

			for j := range monkey.Items {
				monkey.InspectCount += 1

				monkey.ApplyOperation(j, max)

				item := monkey.Items[j]

				if monkey.TestWorryLevel(j) {
					monkeysMap[trueId].Items = append(monkeysMap[trueId].Items, item)
				} else {
					monkeysMap[falseId].Items = append(monkeysMap[falseId].Items, item)
				}
			}

			monkey.Items = []int{}
		}
	}

	var inspects []int

	for _, monkey = range monkeys {
		inspects = append(inspects, monkey.InspectCount)
	}

	common.QuickSort(inspects, 0, len(inspects)-1)
	bestTwo := inspects[(len(inspects))-2:]

	fmt.Println(bestTwo[0] * bestTwo[1])
}
