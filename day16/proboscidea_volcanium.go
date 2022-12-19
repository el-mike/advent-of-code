package day16

import (
	"el-mike/advent-of-code/common"
	"el-mike/advent-of-code/common/ds"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	InputFilename     = "input.txt"
	TestInputFilename = "test_input.txt"
)

func parseLine(line string) *Valve {
	re := regexp.MustCompile(`Valve ([A-Z]{2}).*=(\d+);.*valves (.*)`)

	results := re.FindAllStringSubmatch(line, -1)[0]

	// index "0" contains entire match.
	name := results[1]
	flowRate, err := strconv.Atoi(results[2])
	if err != nil {
		panic(err)
	}

	leadsTo := strings.Split(results[3], ", ")

	return NewValve(name, flowRate, leadsTo)
}

func ProboscideaVolcanium() {
	scanner, err := common.GetFileScanner("./day16/" + TestInputFilename)
	if err != nil {
		panic(err)
	}

	var valvesMap ValvesMap
	var rootValve *Valve

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		valve := parseLine(line)
		valvesMap[valve.Name] = valve

		if rootValve == nil {
			rootValve = valve
		}

		fmt.Println(valve.Name)
	}

	var bestPath *Path

	valvesTree := ds.NewTree[*Valve](ds.NewTreeNode[*Valve](rootValve, nil))

	stepInto := func(
		node *ds.TreeNode[*Valve],
		currentPath *Path,
		minutesLeft int,
	) int {
		minutesLeft -= 1
		if minutesLeft == 0 {
			return 0
		}

	}

	for _, child := range valvesTree.Root {
		stepInto(child, NewP)
	}
}
