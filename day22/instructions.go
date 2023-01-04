package day22

import "strconv"

type DirectionInstruction string

const (
	InstructionLeft  DirectionInstruction = "L"
	InstructionRight                      = "R"
)

type Instructions struct {
	Numbers    []int
	Directions []DirectionInstruction
}

func NewInstructions(line string) *Instructions {
	var numbers []int
	var directions []DirectionInstruction

	currentNumber := ""

	for _, r := range line {
		if DirectionInstruction(r) == InstructionLeft || DirectionInstruction(r) == InstructionRight {
			number, err := strconv.Atoi(currentNumber)
			if err != nil {
				panic(err)
			}

			numbers = append(numbers, number)
			directions = append(directions, DirectionInstruction(r))

			currentNumber = ""
		} else {
			currentNumber += string(r)
		}
	}

	if currentNumber != "" {
		number, err := strconv.Atoi(currentNumber)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, number)
	}

	return &Instructions{
		Numbers:    numbers,
		Directions: directions,
	}
}
