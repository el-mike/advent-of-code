package day_24

import "fmt"

type Vector struct {
	X int
	Y int
}

func NewVector(x, y int) Vector {
	return Vector{
		X: x,
		Y: y,
	}
}

func (v Vector) ID() string {
	return fmt.Sprintf("%d|%d", v.X, v.Y)
}

func (v Vector) Same(candidate Vector) bool {
	return v.X == candidate.X && v.Y == candidate.Y
}

func (v Vector) GetRight() Vector {
	clone := v
	clone.X += 1

	return clone
}

func (v Vector) GetLeft() Vector {
	clone := v
	clone.X -= 1

	return clone
}

func (v Vector) GetUp() Vector {
	clone := v
	clone.Y -= 1

	return clone
}

func (v Vector) GetDown() Vector {
	clone := v
	clone.Y += 1

	return clone
}

func (v Vector) GetNeighbors() []Vector {
	return []Vector{v.GetRight(), v.GetLeft(), v.GetUp(), v.GetDown()}
}
