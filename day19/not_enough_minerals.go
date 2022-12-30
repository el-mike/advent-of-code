package day19

import (
	"el-mike/advent-of-code/common"
	"fmt"
	"sync"
)

const TimeLimit = 32

func NotEnoughMinerals() {
	scanner, err := common.GetFileScanner("./day19/" + common.TestInputFilename)
	if err != nil {
		panic(err)
	}

	var blueprints []*Blueprint

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		numbers, err := common.GetNumbersFromLine(line)
		if err != nil {
			panic(err)
		}

		blueprints = append(blueprints, NewBlueprint(
			numbers[0],
			numbers[1],
			numbers[2],
			[2]int{numbers[3], numbers[4]},
			[2]int{numbers[5], numbers[6]},
		))
	}

	geodesMap := map[int]int{}
	wg := sync.WaitGroup{}

	for _, blueprint := range blueprints {
		wg.Add(1)

		go func(blueprint *Blueprint) {
			startState := NewSimulationState(blueprint)

			geodes := step(blueprint, startState, TimeLimit)
			geodesMap[blueprint.ID] = geodes

			wg.Done()
		}(blueprint)
	}

	wg.Wait()

	totalQuality := 1

	for _, quality := range geodesMap {
		totalQuality *= quality
	}

	fmt.Println(totalQuality)
}

func step(
	blueprint *Blueprint,
	state *SimulationState,
	timeLeft int,
) int {
	timeLeft -= 1

	if timeLeft < 0 {
		return state.Geode
	}

	maxGeodes := 0

	if state.CanBuildOreRobot() &&
		// We only build ore robots when we have less robots than the collective need
		// for ore for one minute - any excessive ore would be wasted at this point.
		// This is one of the heuristics that strongly limits the search space.
		state.OreRobots <= blueprint.OreCostForAllRobots() &&
		state.Ore <= 2*blueprint.OreRobotCost {
		newState := state.Clone()

		// Order is important here! If we update materials after building a robot,
		// new robot would also be included in addition, even though we've just built it.
		newState.UpdateMaterials()
		newState.BuildOreRobot()

		geodes := step(blueprint, newState, timeLeft)

		if geodes > maxGeodes {
			maxGeodes = geodes
		}
	}

	if state.CanBuildClayRobot() &&
		// As above, we only build clay robots when we have less robots
		// than single obsidian robot needs.
		state.ClayRobots <= blueprint.ObsidianRobotCost[1] &&
		state.Ore <= 2*blueprint.ClayRobotCost {
		newState := state.Clone()

		newState.UpdateMaterials()
		newState.BuildClayRobot()

		geodes := step(blueprint, newState, timeLeft)

		if geodes > maxGeodes {
			maxGeodes = geodes
		}
	}

	if state.CanBuildObsidianRobot() &&
		// As above, we only build obsidian robots when we have less robots
		// than single geode robot needs.
		state.ObsidianRobots <= blueprint.GeodeRobotCost[1] {
		newState := state.Clone()

		newState.UpdateMaterials()
		newState.BuildObsidianRobot()

		geodes := step(blueprint, newState, timeLeft)

		if geodes > maxGeodes {
			maxGeodes = geodes
		}
	}

	if state.CanBuildGeodeRobot() {
		newState := state.Clone()

		newState.UpdateMaterials()
		newState.BuildGeodeRobot()

		geodes := step(blueprint, newState, timeLeft)

		if geodes > maxGeodes {
			maxGeodes = geodes
		}
	}

	state.UpdateMaterials()

	// A branch in which no robot is built.
	geodes := step(blueprint, state.Clone(), timeLeft)

	if geodes > maxGeodes {
		maxGeodes = geodes
	}

	return maxGeodes
}
