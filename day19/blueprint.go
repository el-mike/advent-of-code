package day19

type Blueprint struct {
	ID                int
	OreRobotCost      int
	ClayRobotCost     int
	ObsidianRobotCost [2]int
	GeodeRobotCost    [2]int
}

func NewBlueprint(
	id,
	oreCost,
	clayCost int,
	obsidianCost,
	geodeCost [2]int,

) *Blueprint {
	return &Blueprint{
		ID:                id,
		OreRobotCost:      oreCost,
		ClayRobotCost:     clayCost,
		ObsidianRobotCost: obsidianCost,
		GeodeRobotCost:    geodeCost,
	}
}

// OreCostForAllRobots - returns the ore cost for all the robots combined.
// Above this value, there is no point in simulating more ore robots, as any
// excessive ore would be wasted.
func (b *Blueprint) OreCostForAllRobots() int {
	return b.OreRobotCost + b.ClayRobotCost + b.ObsidianRobotCost[0] + b.GeodeRobotCost[0]
}
