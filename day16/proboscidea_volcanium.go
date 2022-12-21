package day16

import (
	"el-mike/advent-of-code/common"
	"el-mike/advent-of-code/common/ds"
	"fmt"
	"log"
	"time"
)

const (
	InputFilename     = "input.txt"
	TestInputFilename = "test_input.txt"
)

const TimeLimit = 30

func ProboscideaVolcanium() {
	start := time.Now()

	scanner, err := common.GetFileScanner("./day16/" + TestInputFilename)
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

	fmt.Println(rootValve.Name)

	queue := ds.NewPriorityQueue[int](func(data []int, i, j int) bool {
		return data[i] > data[j]
	})

	initialValues := []int{9, 11, 18, 13, 15, 14, 7, 8, 12, 10, 4, 6, 3}

	for _, x := range initialValues {
		queue.Enqueue(x)
	}
	
	max, _ := queue.Dequeue()
	fmt.Println(max)

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}
