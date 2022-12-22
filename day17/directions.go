package day17

type Direction uint8

type Directions []Direction

const (
	Left Direction = iota
	Right
	Up
	Down
)

var directionsMap = map[rune]Direction{
	'<': Left,
	'>': Right,
}

// GetTranslateValue - returns a translation value for given axis.
// Please note that due to the board growing upwards, Down direction means
// subtracting from Y axis.
func (d Direction) GetTranslateValue() int {
	if d == Left || d == Down {
		return -1
	} else {
		return 1
	}
}

func (d Direction) IsHorizontal() bool {
	return d == Left || d == Right
}

func (d Direction) IsVertical() bool {
	return d == Up || d == Down
}

func (d *Directions) FromString(line string) {
	for _, r := range line {
		*d = append(*d, directionsMap[r])
	}
}
