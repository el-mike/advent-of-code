package day10

import (
	"el-mike/advent-of-code/common"
	"el-mike/advent-of-code/common/ds"
	"fmt"
	"strconv"
	"strings"
)

func CathodeRayTube() {
	scanner, err := common.GetFileScanner("./day10/" + common.InputFilename)
	if err != nil {
		panic(err)
	}

	operationsQueue := ds.NewQueue[*Operation]()

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")

		var operation *Operation

		if len(parts) == 2 {
			operation = NewOperation(parts[0], parts[1])
		} else {
			operation = NewOperation(parts[0], "")
		}

		operationsQueue.Enqueue(operation)
	}

	clock := NewCpuClock(func() bool {
		return operationsQueue.IsEmpty()
	})

	registerValue := 1

	clock.RegisterCb(func(i int) {
		drawX := (i - 1) % 40

		if drawX == 0 {
			fmt.Print("\n")
		}

		// spritePosition indicates sprite's middle sign.
		spritePosition := registerValue

		if drawX >= (spritePosition-1) && drawX <= (spritePosition+1) {
			fmt.Print("X")
		} else {
			fmt.Print(".")
		}

	})

	executionStack := ds.NewStack[*Operation]()
	executionCounter := 0

	clock.RegisterCb(func(i int) {
		currentOperation, err := executionStack.Peek()
		if _, ok := err.(*ds.StackEmptyException); err != nil && !ok {
			panic(err)
		}

		if currentOperation == nil {
			currentOperation, err = operationsQueue.Dequeue()
			if _, ok := err.(*ds.QueueEmptyException); err != nil && !ok {
				panic(err)
			}

			executionStack.Push(currentOperation)
			executionCounter = currentOperation.ExecutionTime
		}

		executionCounter -= 1

		if executionCounter == 0 {
			if currentOperation != nil && currentOperation.Name == AddOperation {
				value, err := strconv.Atoi(currentOperation.Argument)
				if err != nil {
					panic(err)
				}

				registerValue += value
			}

			_, err := executionStack.Pop()
			if err != nil {
				panic(err)
			}
		}
	})

	clock.Start()
}
