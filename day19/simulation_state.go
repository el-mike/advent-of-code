package day19

type SimulationState struct {
	Blueprint      *Blueprint
	OreRobots      int
	Ore            int
	ClayRobots     int
	Clay           int
	ObsidianRobots int
	Obsidian       int
	GeodeRobots    int
	Geode          int
}

func NewSimulationState(blueprint *Blueprint) *SimulationState {
	return &SimulationState{
		Blueprint:      blueprint,
		OreRobots:      1,
		Ore:            0,
		ClayRobots:     0,
		Clay:           0,
		ObsidianRobots: 0,
		Obsidian:       0,
		GeodeRobots:    0,
		Geode:          0,
	}
}

func (ss *SimulationState) Clone() *SimulationState {
	return &SimulationState{
		Blueprint:      ss.Blueprint,
		OreRobots:      ss.OreRobots,
		Ore:            ss.Ore,
		ClayRobots:     ss.ClayRobots,
		Clay:           ss.Clay,
		ObsidianRobots: ss.ObsidianRobots,
		Obsidian:       ss.Obsidian,
		GeodeRobots:    ss.GeodeRobots,
		Geode:          ss.Geode,
	}
}

func (ss *SimulationState) UpdateMaterials() {
	ss.Ore += ss.OreRobots
	ss.Clay += ss.ClayRobots
	ss.Obsidian += ss.ObsidianRobots
	ss.Geode += ss.GeodeRobots
}

func (ss *SimulationState) CanBuildOreRobot() bool {
	return ss.Ore >= ss.Blueprint.OreRobotCost
}

func (ss *SimulationState) BuildOreRobot() {
	ss.Ore -= ss.Blueprint.OreRobotCost
	ss.OreRobots += 1
}

func (ss *SimulationState) CanBuildClayRobot() bool {
	return ss.Ore >= ss.Blueprint.ClayRobotCost
}

func (ss *SimulationState) BuildClayRobot() {
	ss.Ore -= ss.Blueprint.ClayRobotCost
	ss.ClayRobots += 1
}

func (ss *SimulationState) CanBuildObsidianRobot() bool {
	return ss.Ore >= ss.Blueprint.ObsidianRobotCost[0] &&
		ss.Clay >= ss.Blueprint.ObsidianRobotCost[1]
}

func (ss *SimulationState) BuildObsidianRobot() {
	ss.Ore -= ss.Blueprint.ObsidianRobotCost[0]
	ss.Clay -= ss.Blueprint.ObsidianRobotCost[1]
	ss.ObsidianRobots += 1
}

func (ss *SimulationState) CanBuildGeodeRobot() bool {
	return ss.Ore >= ss.Blueprint.GeodeRobotCost[0] &&
		ss.Obsidian >= ss.Blueprint.GeodeRobotCost[1]
}

func (ss *SimulationState) BuildGeodeRobot() {
	ss.Ore -= ss.Blueprint.GeodeRobotCost[0]
	ss.Obsidian -= ss.Blueprint.GeodeRobotCost[1]
	ss.GeodeRobots += 1
}
