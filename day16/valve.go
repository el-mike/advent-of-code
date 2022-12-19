package day16

type ValvesMap map[string]*Valve

type Valve struct {
	Name     string
	FlowRate int
	LeadsTo  []string
}

func NewValve(name string, flowRate int, leadsTo []string) *Valve {
	return &Valve{
		Name:     name,
		FlowRate: flowRate,
		LeadsTo:  leadsTo,
	}
}
