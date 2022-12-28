package day18

import (
	"el-mike/advent-of-code/common"
	"el-mike/advent-of-code/common/ds"
	"fmt"
	"strconv"
	"strings"
)

func BoilingBoulders() {
	scanner, err := common.GetFileScanner("./day18/" + common.InputFilename)
	if err != nil {
		panic(err)
	}

	var cubes []*Cube

	maxIndex := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		var coords Coords

		for i, numberStr := range strings.Split(line, ",") {
			value, err := strconv.Atoi(numberStr)
			if err != nil {
				panic(err)
			}

			// We add one, as we want to create a "padding" for the matrix, so every axis
			// has a one row of space. It will help us with finding exterior surface.
			coords[i] = value + 1

			if coords[i] > maxIndex {
				maxIndex = coords[i]
			}
		}

		cubes = append(cubes, NewCube(coords))
	}

	// This time, we need to add 2 - one to accommodate for changing from index to size,
	// and one to add another "padding", this time as the end.
	matrixSize := maxIndex + 2

	matrixModel := NewMatrixModel(matrixSize)

	for x := 0; x < matrixSize; x += 1 {
		matrixModel.Matrix[x] = make([][]MatrixCell, matrixSize)

		for y := 0; y < matrixSize; y += 1 {
			matrixModel.Matrix[x][y] = make([]MatrixCell, matrixSize)

			for z := 0; z < matrixSize; z += 1 {
				matrixModel.Matrix[x][y][z] = EmptyCell
			}
		}
	}

	for _, cube := range cubes {
		x, y, z := cube.Coords[0], cube.Coords[1], cube.Coords[2]

		matrixModel.Matrix[x][y][z] = LavaCell
	}

	externalSurfaceArea := FindExteriorSurfaceArea(matrixModel)

	fmt.Println(externalSurfaceArea)
}

// FindExteriorSurfaceArea - 3D Breadth-first search algorithm. It checks every Empty cell
// in the matrix, and adds 1 to totalSurface if it detecs a LavaCell neighbor.
func FindExteriorSurfaceArea(matrixModel *MatrixModel) int {
	totalSurface := 0

	frontier := ds.NewQueue[Coords]()
	frontier.Enqueue(Coords{0, 0, 0})

	for !frontier.IsEmpty() {
		current, err := frontier.Dequeue()
		if err != nil {
			panic(err)
		}

		neighbors := matrixModel.GetNeighbors(current)

		for _, neighbor := range neighbors {
			cell := matrixModel.At(neighbor)

			if cell == LavaCell {
				totalSurface += 1
			}
			if cell == EmptyCell {
				frontier.Enqueue(neighbor)
				matrixModel.SetAt(neighbor, SteamCell)
			}
		}

	}

	return totalSurface
}
