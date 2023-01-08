package day16

import (
	"el-mike/advent-of-code/go/common"
	"fmt"
	"log"
	"sync"
	"time"
)

type ResultWrapper struct {
	Max      int
	bestPath *Path
	PathStr  string
}

const TimeLimit = 30

func ProboscideaVolcanium() {
	start := time.Now()

	scanner, err := common.GetFileScanner("./year_2022/day_16/" + common.InputFilename)
	if err != nil {
		panic(err)
	}

	parser := NewParser()

	valvesMap := ValvesMap{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		valve := parser.ParseLine(line)
		valvesMap[valve.Name] = valve
	}

	rootValve := valvesMap["AA"]
	numValves := 0

	for _, valve := range valvesMap {
		if valve.FlowRate >= 0 {
			numValves += 1
		}
	}

	resultWrapper := &ResultWrapper{
		bestPath: NewPath(numValves),
	}

	wg := sync.WaitGroup{}

	for _, valveName := range rootValve.LeadsTo {
		wg.Add(1)

		go func(valveName string) {
			path := NewPath(numValves)

			stepInto(TimeLimit, valvesMap[valveName], path, rootValve.Name, resultWrapper, valvesMap)
			wg.Done()

		}(valveName)
	}

	wg.Wait()

	fmt.Println(resultWrapper.bestPath.Total)

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func stepInto(
	minutesLeft int,
	currentValve *Valve,
	currentPath *Path,
	cameFrom string,
	resultWrapper *ResultWrapper,
	valvesMap ValvesMap,
) {
	minutesLeft -= 1

	if minutesLeft == 0 || currentPath.AllValvesOpened() {
		if currentPath.Total > resultWrapper.bestPath.Total {
			resultWrapper.bestPath = currentPath
		}

		return
	}

	for _, valveName := range currentValve.LeadsTo {
		if valveName == cameFrom {
			continue
		}

		stepInto(minutesLeft, valvesMap[valveName], currentPath.Clone(), currentValve.Name, resultWrapper, valvesMap)
	}

	if currentValve.FlowRate != 0 && !currentPath.HasBeenOpened(currentValve.Name) {
		minutesLeft -= 1

		if minutesLeft == 0 {
			return
		}

		currentPath.Open(currentValve.Name)
		currentPath.Total += currentValve.FlowRate * minutesLeft

		for _, valveName := range currentValve.LeadsTo {
			stepInto(minutesLeft, valvesMap[valveName], currentPath.Clone(), currentValve.Name, resultWrapper, valvesMap)
		}
	}
}
