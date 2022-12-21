package day16

import (
	"el-mike/advent-of-code/common/ds"
	"math"
)

type DistancesMap map[string]int
type PreviousMap map[string]string
type VisitedMap map[string]bool

type PathResult struct {
	Steps []string
	Cost  int
}

func getPriority(item *ds.Vertex[*Valve]) int {
	return item.Weight
}

func Dijkstra(
	graph *ds.WeightedGraph[*Valve],
	sourceVertex *ds.Vertex[*Valve],
	targetVertexId string,
) *PathResult {
	// Using min priority queue improves performance, as we want to prioritize the
	// edges with minimal weight.
	queue := ds.NewMinPriorityQueue[*ds.Vertex[*Valve]](getPriority)

	// Stores the distances from the start vertex to every other vertex.
	distances := DistancesMap{}
	// Stores previous ID as value for every ID as a key.
	previous := PreviousMap{}
	// Stores information about vertices that have already been checked.
	visited := VisitedMap{}

	// We mark every vertex as unvisited, and set its distance from starting vertex
	// to MaxInt (there is no Int infinity in Go).
	// We set starting vertex's distance to 0, as it's used in consecutive calculations.
	for id, _ := range graph.AdjacencyList {
		visited[id] = false

		if id == sourceVertex.ID {
			distances[id] = 0
			continue
		}

		distances[id] = math.MaxInt
	}

	// Starting vertex has to have Weight 0 (same as distance), so in order to make sure,
	// we clone it and set the weight.
	start := sourceVertex.Clone()
	start.Weight = 0

	queue.Enqueue(start)

	for !queue.IsEmpty() {
		current, _ := queue.Dequeue()

		// For every neighbor that we haven't visited yet, check its distance from the START VERTEX.
		// If it's smaller, replace the distance and the node that we arrived from at given neighbor.
		// Enqueue neighbor to be checked in upcoming iterations.
		for _, neighbor := range graph.AdjacencyList[current.ID] {
			if visited[neighbor.ID] == true {
				continue
			}

			distance := distances[current.ID] + neighbor.Weight

			relativeNeighbor := neighbor.Clone()
			relativeNeighbor.Weight = distance

			queue.Enqueue(relativeNeighbor)

			if distance < distances[neighbor.ID] {
				distances[neighbor.ID] = distance
				previous[neighbor.ID] = current.ID
			}
		}

		visited[current.ID] = true
	}

	pathStep := targetVertexId

	result := &PathResult{
		Cost:  distances[targetVertexId],
		Steps: []string{targetVertexId},
	}

	for pathStep != sourceVertex.ID {
		pathStep = previous[pathStep]
		result.Steps = append(result.Steps, pathStep)
	}

	return result
}
