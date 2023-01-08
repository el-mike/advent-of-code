package day_05

import (
	"errors"
	"regexp"
	"strconv"
)

type Operation struct {
	Origin string

	Amount int
	From   int
	To     int
}

func NewOperation(operation string) *Operation {
	amount, from, to := parseOperation(operation)

	return &Operation{
		Origin: operation,
		Amount: amount,
		From:   from,
		To:     to,
	}
}

func parseOperation(operation string) (int, int, int) {
	re := regexp.MustCompile("[0-9]+")

	results := re.FindAllString(operation, -1)

	// Every operation should contain 3 numbers - amount of crates to move, index of source
	// stack and index of target stack.
	if len(results) != 3 {
		panic(errors.New("Operation malformed"))
	}

	amount, err := strconv.Atoi(results[0])
	if err != nil {
		panic(err)
	}

	from, err := strconv.Atoi(results[1])
	if err != nil {
		panic(err)
	}

	to, err := strconv.Atoi(results[2])
	if err != nil {
		panic(err)
	}

	return amount, from, to
}
