package day18

type MatrixCell uint8

const (
	EmptyCell MatrixCell = iota
	LavaCell
	SteamCell
)

type MatrixModel struct {
	Matrix [][][]MatrixCell
	Size   int
}

func NewMatrixModel(size int) *MatrixModel {
	return &MatrixModel{
		Matrix: make([][][]MatrixCell, size),
		Size:   size,
	}
}

func (m *MatrixModel) At(coords Coords) MatrixCell {
	return m.Matrix[coords[0]][coords[1]][coords[2]]
}

func (m *MatrixModel) SetAt(coords Coords, value MatrixCell) {
	m.Matrix[coords[0]][coords[1]][coords[2]] = value
}

func (m *MatrixModel) GetNeighbors(target Coords) []Coords {
	x, y, z := target[0], target[1], target[2]

	x1, x2 := x-1, x+1
	y1, y2 := y-1, y+1
	z1, z2 := z-1, z+1

	var neighbors []Coords

	if x1 >= 0 {
		neighbors = append(neighbors, Coords{x1, y, z})
	}
	if x2 < m.Size {
		neighbors = append(neighbors, Coords{x2, y, z})
	}
	if y1 >= 0 {
		neighbors = append(neighbors, Coords{x, y1, z})
	}
	if y2 < m.Size {
		neighbors = append(neighbors, Coords{x, y2, z})
	}
	if z1 >= 0 {
		neighbors = append(neighbors, Coords{x, y, z1})
	}
	if z2 < m.Size {
		neighbors = append(neighbors, Coords{x, y, z2})
	}

	return neighbors
}
