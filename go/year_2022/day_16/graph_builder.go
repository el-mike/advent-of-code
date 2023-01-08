package day_16

import (
	"el-mike/advent-of-code/go/common/ds"
)

type NeighborsResultWrapper struct {
	Neighbors []*ds.Vertex[*Valve]
}

type GraphBuilder struct {
	valvesMap ValvesMap
}

func NewGraphBuilder(valvesMap ValvesMap) *GraphBuilder {
	return &GraphBuilder{
		valvesMap: valvesMap,
	}
}

func (gb *GraphBuilder) Build() *ds.WeightedGraph[*Valve] {
	graph := ds.NewWeightedGraph[*Valve]()

	for _, valve := range gb.valvesMap {
		// We want to build a graph from valid valves only, therefore we skip
		// the one with FlowRate zero as "base" vertices.
		if valve.FlowRate == 0 && valve.Name != "AA" {
			continue
		}

		vertex := &ds.Vertex[*Valve]{ID: valve.Name, Data: valve}
		graph.AddVertex(vertex)

		for _, leadsTo := range valve.LeadsTo {
			closestNeighborsWrapper := &NeighborsResultWrapper{
				Neighbors: []*ds.Vertex[*Valve]{},
			}

			gb.findNeighbors(closestNeighborsWrapper, valve.Name, leadsTo, 0)

			for _, neighbor := range closestNeighborsWrapper.Neighbors {
				graph.AddVertex(neighbor)
				// We reuse neighbor.Weight, as it already been determined
				// in the recursive process.
				graph.AddEdge(vertex, neighbor, neighbor.Weight)
			}
		}
	}

	return graph
}

func (gb *GraphBuilder) findNeighbors(
	wrapper *NeighborsResultWrapper,
	previousValveName,
	currentValveName string,
	weight int,
) {
	weight += 1

	currentValve := gb.valvesMap[currentValveName]

	if currentValve.FlowRate > 0 || currentValveName == "AA" {
		wrapper.Neighbors = append(wrapper.Neighbors, &ds.Vertex[*Valve]{
			ID:     currentValve.Name,
			Data:   currentValve,
			Weight: weight,
		})

		return
	}

	for _, childValveName := range currentValve.LeadsTo {
		// This guard prevents circular motion.
		if childValveName == previousValveName {
			continue
		}

		gb.findNeighbors(wrapper, currentValveName, childValveName, weight)
	}
}
