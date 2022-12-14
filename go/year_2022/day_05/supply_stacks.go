package day_05

import (
	"el-mike/advent-of-code/go/common"
	"el-mike/advent-of-code/go/common/ds"
	"errors"
	"math"
	"strings"
)

const (
	COLUMN_SIZE = 4
)

// Returns initial crates state in a form of slice of stacks containing symbols.
// Each stack has the same order as defined in the input.
func parseInitialState(initialState *ds.Stack[string], numColumns int) []*ds.Stack[string] {
	stacks := make([]*ds.Stack[string], numColumns)

	for !initialState.Empty() {
		line, err := initialState.Pop()
		if err != nil {
			panic(err)
		}

		for i := 0; i < numColumns; i += 1 {
			if stacks[i] == nil {
				stacks[i] = ds.NewStack[string]()
			}

			// Each column is 4 characters wide.
			start := i * COLUMN_SIZE
			end := start + COLUMN_SIZE

			// Last column does not have right separator, therefore we use the length of
			// the string.
			if i == (numColumns - 1) {
				end = len(line)
			}

			cell := line[start:end]

			// If cell does not have a crate symbol inside, we simply skip it.
			if !strings.Contains(cell, "[") {
				continue
			}

			symbol := strings.Trim(cell, " ")
			symbol = strings.ReplaceAll(symbol, "[", "")
			symbol = strings.ReplaceAll(symbol, "]", "")

			stacks[i].Push(symbol)
		}
	}

	return stacks
}

func SupplyStacks() string {
	scanner, err := common.GetFileScanner("./year_2022/day_05/" + common.InputFilename)
	if err != nil {
		panic(err)
	}

	initialState := ds.NewStack[string]()

	for scanner.Scan() {
		line := scanner.Text()

		// If line equals zero, that means we reached the end of the starting crates positions.
		if line == "" {
			break
		} else {
			initialState.Push(line)
		}
	}

	// We want to discard the first row, as it's going to contain only column numbers
	// from the bottom.
	initialState.Pop()

	row, err := initialState.Peek()
	if err != nil {
		panic(err)
	}

	numColumns := int(math.Ceil(float64(len(row)) / COLUMN_SIZE))

	stacks := parseInitialState(initialState, numColumns)

	// After parsing the initial state, we can resume scanning - it will start with the first
	// move operation description.
	for scanner.Scan() {
		line := scanner.Text()

		// Guard for reading empty lines.
		if line == "" {
			continue
		}

		operation := NewOperation(line)

		// We need to subtract the one from the From and To, as indexes are zero-based.
		source := stacks[operation.From-1]
		target := stacks[operation.To-1]

		if operation.Amount > source.Size() {
			panic(errors.New("amount to move is bigger than stack's size"))
		}

		// The simplest way to retain the order between two stacks is to introduce
		// additional, temporary stack, which will have the reversed order of elements.
		tmpStack := ds.NewStack[string]()

		for i := 0; i < operation.Amount; i += 1 {
			item, err := source.Pop()
			if err != nil {
				panic(err)
			}

			tmpStack.Push(item)
		}

		// We need to access tmpSize size, as using Stack() in loop declaration
		// would return different number with each iteration.
		tmpSize := tmpStack.Size()

		for i := 0; i < tmpSize; i += 1 {
			item, err := tmpStack.Pop()
			if err != nil {
				panic(err)
			}

			target.Push(item)
		}
	}

	result := ""

	for i := 0; i < numColumns; i += 1 {
		item, err := stacks[i].Peek()
		if err != nil {
			panic(err)
		}

		result += item
	}

	return result
}
