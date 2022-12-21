package day16

import "el-mike/advent-of-code/common"

type Path struct {
	Traversed []string
	Opened    []string
	NumValves int
	Total     int
}

func NewPath(numValves int) *Path {
	return &Path{
		Traversed: []string{},
		Opened:    []string{},
		NumValves: numValves,
		Total:     0,
	}
}

func (p *Path) AddStep(name string) {
	p.Traversed = append(p.Traversed, name)
}

func (p *Path) Open(name string) {
	p.Opened = append(p.Opened, name)
}

func (p *Path) AllValvesOpened() bool {
	return len(p.Opened) == p.NumValves
}

func (p *Path) HasBeenOpened(name string) bool {
	return common.Contains[string](p.Opened, name)
}

func (p *Path) AddTotal(pressureReleased int) {
	p.Total += pressureReleased
}

func (p *Path) Clone() *Path {
	return &Path{
		Traversed: append([]string{}, p.Traversed...),
		Opened:    append([]string{}, p.Opened...),
		NumValves: p.NumValves,
		Total:     p.Total,
	}
}
