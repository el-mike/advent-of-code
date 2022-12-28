package day18

type Coords [3]int

func (c Coords) Same(candidate Coords) bool {
	return c[0] == candidate[0] && c[1] == candidate[1] && c[2] == candidate[2]
}
