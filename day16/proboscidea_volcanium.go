package day16

import (
	"el-mike/advent-of-code/common"
	"el-mike/advent-of-code/common/ds"
	"fmt"
	"log"
	"time"
)

type PathsMap map[string]map[string]int
type OpenedMap map[string]bool

type Toolbox struct {
	BestPath  *Path
	ValvesMap ValvesMap
	PathsMap  PathsMap
}

const (
	InputFilename     = "input.txt"
	TestInputFilename = "test_input.txt"
)

const TimeLimit = 30
const StartValveName = "AA"

func ProboscideaVolcanium() {
	start := time.Now()

	scanner, err := common.GetFileScanner("./day16/" + TestInputFilename)
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

	pathsMap := PathsMap{}

	for sourceId, _ := range graph.AdjacencyList {
		pathsMap[sourceId] = map[string]int{}

		for targetId, _ := range graph.AdjacencyList {
			if sourceId == targetId {
				continue
			}

			source := &ds.Vertex[*Valve]{ID: sourceId, Data: valvesMap[sourceId], Weight: 0}
			path := Dijkstra(graph, source, targetId)

			pathsMap[sourceId][targetId] = path.Cost
		}
	}

	toolbox := &Toolbox{
		BestPath:  NewPath(numValves),
		ValvesMap: valvesMap,
		PathsMap:  pathsMap,
	}

	startPaths := pathsMap[StartValveName]

	for targetId, _ := range startPaths {
		path := NewPath(numValves)
		path.AddStep(StartValveName)

		step(toolbox, TimeLimit, StartValveName, targetId, path)
	}

	fmt.Println(toolbox.BestPath.Total)

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func step(
	toolbox *Toolbox,
	timeLeft int,
	cameFrom string,
	currentId string,
	currentPath *Path,
) {
	currentPath.AddStep(currentId)

	// As we always open a valve upon a single step, we add "1" as valve opening time.
	timeSpent := toolbox.PathsMap[cameFrom][currentId] + 1
	timeLeft -= timeSpent

	if timeLeft <= 0 {
		if currentPath.Total > toolbox.BestPath.Total {
			toolbox.BestPath = currentPath
		}

		return
	}

	currentPath.Open(currentId)
	currentPath.Total += toolbox.ValvesMap[currentId].FlowRate * timeLeft

	hasCandidatesLeft := false

	for candidateId, _ := range toolbox.PathsMap[currentId] {
		if candidateId != StartValveName && !currentPath.HasBeenOpened(candidateId) {
			hasCandidatesLeft = true
			step(toolbox, timeLeft, currentId, candidateId, currentPath.Clone())
		}
	}

	if !hasCandidatesLeft {
		if currentPath.Total > toolbox.BestPath.Total {
			toolbox.BestPath = currentPath
		}
	}
}
