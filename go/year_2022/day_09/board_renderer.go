package day_09

import "fmt"

type BoardRenderer struct {
	NumCols int
	NumRows int
	Start   []int
}

func NewBoardRenderer() *BoardRenderer {
	return &BoardRenderer{
		NumCols: 30,
		NumRows: 25,
		Start:   []int{15, 15},
	}
}

func (br *BoardRenderer) render(positions [][]int) {
	for i := 0; i < br.NumRows; i += 1 {
		for j := 0; j < br.NumCols; j += 1 {
			sign := ""

			for k := len(positions) - 1; k >= 0; k -= 1 {
				x := br.Start[0] + positions[k][0]
				y := br.Start[1] + positions[k][1]

				if x == i && y == j {
					if k == 0 {
						sign = "H"
					} else {
						sign = fmt.Sprintf("%d", k)
					}
				}
			}

			if sign != "" {
				fmt.Printf(sign)
			} else {
				fmt.Printf(".")
			}
		}

		fmt.Printf("\n")
	}

	fmt.Printf("\n\n")
}
