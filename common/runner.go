package common

import (
	"log"
	"time"
)

type Runnable func()

type Runner struct{}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) RunAndMeasure(runnable Runnable) {
	start := time.Now()

	runnable()

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}
