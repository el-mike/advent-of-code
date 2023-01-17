package day_24

type Coord struct {
	X int
	Y int
}

func NewCoord(x, y int) Coord {
	return Coord{
		X: x,
		Y: y,
	}
}

func (c Coord) Same(candidate Coord) bool {
	return c.X == candidate.X && c.Y == candidate.Y
}

func (c Coord) GetRight() Coord {
	clone := c
	clone.X += 1

	return clone
}

func (c Coord) GetLeft() Coord {
	clone := c
	clone.X -= 1

	return clone
}

func (c Coord) GetUp() Coord {
	clone := c
	clone.Y -= 1

	return clone
}

func (c Coord) GetDown() Coord {
	clone := c
	clone.Y += 1

	return clone
}
