package day_24

import "fmt"

type StepInfo struct {
	Minute   int
	Position Vector
}

func (si *StepInfo) ID() string {
	return fmt.Sprintf("%d|%s", si.Minute, si.Position.ID())
}
