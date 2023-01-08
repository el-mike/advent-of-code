package day_11

import (
	"errors"
	"strconv"
)

const (
	OldOperand = "old"
)

type OperationCb func(worryLevel, operand int) int

var OperationsMap = map[string]OperationCb{
	"+": func(worryLevel, operand int) int { return worryLevel + operand },
	"*": func(worryLevel, operand int) int { return worryLevel * operand },
}

type Monkey struct {
	ID           int
	Items        []int
	Operation    OperationCb
	Operand      string
	TestValue    int
	Recipients   []int
	InspectCount int
}

func NewMonkey() *Monkey {
	return &Monkey{
		Items:        []int{},
		InspectCount: 0,
	}
}

func (m *Monkey) ApplyOperation(itemIndex, max int) {
	if itemIndex < 0 || itemIndex >= len(m.Items) {
		panic(errors.New("index outside bounds"))
	}

	currentValue := m.Items[itemIndex]
	numOperand := 0

	if m.Operand == OldOperand {
		numOperand = currentValue
	} else {
		number, err := strconv.Atoi(m.Operand)
		if err != nil {
			panic(err)
		}

		numOperand = number
	}

	m.Items[itemIndex] = m.Operation(currentValue, numOperand) % max
}

func (m *Monkey) TestWorryLevel(itemIndex int) bool {
	return (m.Items[itemIndex] % m.TestValue) == 0
}
