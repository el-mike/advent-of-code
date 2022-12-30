package day19

import (
	"el-mike/advent-of-code/common"
	"fmt"
	"sync"
)

const TimeLimit = 32

func NotEnoughMinerals() {
	scanner, err := common.GetFileScanner("./day19/" + common.InputFilename)
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

	wg := sync.WaitGroup{}
	result := 1

	for _, blueprint := range blueprints[:3] {
		wg.Add(1)

		go func(blueprint *Blueprint) {
			startState := NewSimulationState(blueprint)

			result *= step(blueprint, startState, TimeLimit)

			wg.Done()
		}(blueprint)
	}

	wg.Wait()
	
	fmt.Println(result)
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
		// If there is a lot of ore and nothing is being built, we don't want to pursue given path.
		state.Ore <= 2*blueprint.OreRobotCost &&
		// If there is only 10 minutes left we should already by collecting enough ore.
		timeLeft > 10 {
		newState := state.Clone()

		// Order is important here! If we update materials after building a robot,
		// new robot would also be included in addition, even though we've just built it.
		// Also, we cannot update materials before testing if given robot can be built, as it would
		// allow to build robots too early (even if given round ends up with enough material,
		// robots are always built BEFORE collecting minerals, not after).
		newState.UpdateMaterials()
		newState.BuildOreRobot()

		if geodes := step(blueprint, newState, timeLeft); geodes > maxGeodes {
			maxGeodes = geodes
		}
	}

	if state.CanBuildClayRobot() &&
		// As above, we only build clay robots when we have less robots
		// than single obsidian robot needs.
		state.ClayRobots <= blueprint.ObsidianRobotCost[1] &&
		state.Ore <= 2*blueprint.ClayRobotCost &&
		timeLeft > 10 {
		newState := state.Clone()

		newState.UpdateMaterials()
		newState.BuildClayRobot()

		if geodes := step(blueprint, newState, timeLeft); geodes > maxGeodes {
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

		if geodes := step(blueprint, newState, timeLeft); geodes > maxGeodes {
			maxGeodes = geodes
		}
	}

	if state.CanBuildGeodeRobot() {
		newState := state.Clone()

		newState.UpdateMaterials()
		newState.BuildGeodeRobot()

		if geodes := step(blueprint, newState, timeLeft); geodes > maxGeodes {
			maxGeodes = geodes
		}
	}

	state.UpdateMaterials()

	// A branch in which no robot is built. We only want to pursuit this branch if we can't
	// build some robot now, but we could in the future. Otherwise, it's just a waste of time.
	// This makes a huge difference in terms of time complexity.
	if !state.CanBuildOreRobot() ||
		!state.CanBuildClayRobot() ||
		!state.CanBuildObsidianRobot() ||
		!state.CanBuildGeodeRobot() {
		if geodes := step(blueprint, state, timeLeft); geodes > maxGeodes {
			maxGeodes = geodes
		}
	}

	return maxGeodes
}
