package day17

type Coord [2]int

type Coords []Coord

func (c Coords) Clone() Coords {
	cs := append(Coords{}, c...)

	return cs
}

func (c Coords) TranslateX(diffX int) {
	for i := range c {
		c[i][0] += diffX
	}
}

func (c Coords) TranslateY(diffY int) {
	for i := range c {
		c[i][1] += diffY
	}
}

func (c Coords) GetMaxY() int {
	max := 0
	for i := range c {
		y := c[i][1]

		if y > max {
			max = y
		}
	}

	return max
}
