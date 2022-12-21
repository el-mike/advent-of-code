package day16

import "el-mike/advent-of-code/common"

type Path struct {
	Traversed []string
	Opened    []string
	Total     int
}

func NewPath() *Path {
	return &Path{
		Traversed: []string{},
		Opened:    []string{},
		Total:     0,
	}
}

func (p *Path) AddStep(name string) {
	p.Traversed = append(p.Traversed, name)
}

func (p *Path) Open(name string) {
	p.Opened = append(p.Opened, name)
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
		Total:     p.Total,
	}
}
