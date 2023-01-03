package day21

import "math"

type OperationFn func(left, right int) int

const (
	Addition       = "+"
	Subtration     = "-"
	Multiplication = "*"
	Division       = "/"
)

var operationsMap = map[string]OperationFn{
	Addition:       func(left, right int) int { return left + right },
	Subtration:     func(left, right int) int { return left - right },
	Multiplication: func(left, right int) int { return left * right },
	Division:       func(left, right int) int { return left / right },
}

type Monkey struct {
	Name        string
	Number      int
	LeftMonkey  string
	RightMonkey string
	OperationFn OperationFn
}

func NewMonkey(
	name string,
	number int,
	leftMonkey,
	rightMonkey,
	operation string,
) *Monkey {
	var operationFn OperationFn = nil

	if operation != "" {
		operationFn = operationsMap[operation]
	}

	return &Monkey{
		Name:        name,
		Number:      number,
		LeftMonkey:  leftMonkey,
		RightMonkey: rightMonkey,
		OperationFn: operationFn,
	}
}

func (m *Monkey) HasNumber() bool {
	return m.Number != math.MaxInt
}
