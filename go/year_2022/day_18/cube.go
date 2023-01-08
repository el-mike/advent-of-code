package day_18

type Cube struct {
	Coords Coords
}

func NewCube(coords Coords) *Cube {
	return &Cube{
		Coords: coords,
	}
}
