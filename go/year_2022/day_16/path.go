package day_16

import "el-mike/advent-of-code/go/common"

type Path struct {
	Opened    []string
	NumValves int
	Total     int
}

func NewPath(numValves int) *Path {
	return &Path{
		Opened:    []string{},
		NumValves: numValves,
		Total:     0,
	}
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

func (p *Path) Clone() *Path {
	return &Path{
		Opened:    append([]string{}, p.Opened...),
		NumValves: p.NumValves,
		Total:     p.Total,
	}
}
