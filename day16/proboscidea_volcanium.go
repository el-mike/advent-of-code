package day16

import (
	"el-mike/advent-of-code/common"
	"el-mike/advent-of-code/common/ds"
	"fmt"
)

type CostsMap map[string]map[string]int

type HelperContext struct {
	NumValves int
	BestPath  *Path
	ValvesMap ValvesMap
	CostsMap  CostsMap
}

const TimeLimit = 26
const StartValveName = "AA"

func ProboscideaVolcanium() {
	scanner, err := common.GetFileScanner("./day16/" + common.InputFilename)
	if err != nil {
		panic(err)
	}

	parser := NewParser()
	valvesMap := ValvesMap{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		valve := parser.ParseLine(line)
		valvesMap[valve.Name] = valve
	}

	graphBuilder := NewGraphBuilder(valvesMap)
	graph := graphBuilder.Build()

	// We subtract one because of "AA" valve, which needs to be part of the graph,
	// but cannot be opened.
	numValves := len(graph.AdjacencyList) - 1

	costsMap := CostsMap{}

	for sourceId := range graph.AdjacencyList {
		costsMap[sourceId] = map[string]int{}

		for targetId := range graph.AdjacencyList {
			if sourceId == targetId {
				continue
			}

			source := &ds.Vertex[*Valve]{ID: sourceId, Data: valvesMap[sourceId], Weight: 0}
			path := Dijkstra(graph, source, targetId)

			costsMap[sourceId][targetId] = path.Cost
		}
	}

	context := &HelperContext{
		NumValves: numValves,
		BestPath:  NewPath(numValves),
		ValvesMap: valvesMap,
		CostsMap:  costsMap,
	}

	candidates := context.CostsMap[StartValveName]

	for candidateId := range candidates {
		path := NewPath(context.NumValves)
		step(context, TimeLimit, StartValveName, candidateId, path, 1)
	}

	fmt.Println(context.BestPath.Total)
}

func step(
	context *HelperContext,
	timeLeft int,
	cameFrom string,
	currentId string,
	currentPath *Path,
	playersLeft int,
) {
	// As we always open a valve upon a single step, we add "1" as valve opening time.
	timeSpent := context.CostsMap[cameFrom][currentId] + 1
	timeLeft -= timeSpent

	if timeLeft <= 0 {
		if currentPath.Total > context.BestPath.Total {
			context.BestPath = currentPath
		}

		return
	}

	currentPath.Open(currentId)
	currentPath.Total += context.ValvesMap[currentId].FlowRate * timeLeft

	// A bit of heuristic touch - we can assume that most efficient paths distribution will be
	// around 50/50, therefore we only simulate second player if the first one already explored
	// some paths, but not too many.
	if playersLeft > 0 && len(currentPath.Opened) > 5 && len(currentPath.Opened) < 12 {
		// For given state (currentPath) run simulation for the second player.
		for candidateId := range context.CostsMap[StartValveName] {
			if candidateId != StartValveName && !currentPath.HasBeenOpened(candidateId) {
				path := currentPath.Clone()
				step(context, TimeLimit, StartValveName, candidateId, path, playersLeft-1)
			}
		}
	}

	hasCandidatesLeft := false

	for candidateId := range context.CostsMap[currentId] {
		if candidateId != StartValveName && !currentPath.HasBeenOpened(candidateId) {
			hasCandidatesLeft = true
			step(context, timeLeft, currentId, candidateId, currentPath.Clone(), playersLeft)
		}
	}

	if !hasCandidatesLeft {
		if currentPath.Total > context.BestPath.Total {
			context.BestPath = currentPath
		}
	}
}
