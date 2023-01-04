package day22

type Direction uint8

const (
	DirectionRight Direction = iota
	DirectionDown
	DirectionLeft
	DirectionUp
)

func (d Direction) Rotate(instruction DirectionInstruction) Direction {
	if instruction == InstructionLeft {
		if d == DirectionRight {
			return DirectionUp
		} else {
			return d - 1
		}
	} else if instruction == InstructionRight {
		if d == DirectionUp {
			return DirectionRight
		} else {
			return d + 1
		}
	}

	return d
}

func (d Direction) IsHorizontal() bool {
	return d == DirectionRight || d == DirectionLeft
}

func (d Direction) IsVertical() bool {
	return d == DirectionUp || d == DirectionDown
}
