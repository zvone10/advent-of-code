package internal

import (
	"fmt"
	"time"
)

type Function func()

type DayRunner interface {
	Execute()
}

type SequentialDayRunner struct {
	start time.Time
	f1    Function
	f2    Function
}

func NewDayRunner(f1, f2 Function) *SequentialDayRunner {
	return &SequentialDayRunner{
		start: time.Now(),
		f1:    f1,
		f2:    f2,
	}
}

func (s *SequentialDayRunner) Execute() {
	//var elapsed time.Duration
	s.start = time.Now()

	s.f1()
	elapsedP1 := time.Since(s.start)
	fmt.Printf("Part 1 execution time: %d ms\n", elapsedP1.Milliseconds())

	s.f2()
	elapsedP2 := time.Since(s.start)
	fmt.Printf("Part 2 execution time: %d ms\n", (elapsedP2 - elapsedP1).Milliseconds())

	fmt.Printf("Total time: %d ms\n", elapsedP2.Milliseconds())
}
