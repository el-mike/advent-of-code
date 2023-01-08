package day_12

import "el-mike/advent-of-code/go/common/ds"

type StepsMap map[string]*GridCell

func BreadthFirstSearch(model *GridModel, start *GridCell) StepsMap {
	frontier := ds.NewQueue[*GridCell]()
	stepsMap := StepsMap{}

	frontier.Enqueue(start)

	for !frontier.IsEmpty() {
		current, err := frontier.Dequeue()
		if err != nil {
			panic(err)
		}

		neighbors := model.GetNeighbors(current)

		for _, neighbor := range neighbors {
			if stepsMap[neighbor.ID()] == nil {
				frontier.Enqueue(neighbor)
				stepsMap[neighbor.ID()] = current
			}
		}
	}

	return stepsMap
}
