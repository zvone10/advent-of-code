package main

import (
	"flag"
	"fmt"

	"github.com/zvone10/advent-of-code/internal"
	aoc2025d1p1 "github.com/zvone10/advent-of-code/pkg/aoc2025/d1/p1"
	aoc2025d1p2 "github.com/zvone10/advent-of-code/pkg/aoc2025/d1/p2"
	aoc2025d10p1 "github.com/zvone10/advent-of-code/pkg/aoc2025/d10/p1"
	aoc2025d10p2 "github.com/zvone10/advent-of-code/pkg/aoc2025/d10/p2"
	aoc2025d11p1 "github.com/zvone10/advent-of-code/pkg/aoc2025/d11/p1"
	aoc2025d11p2 "github.com/zvone10/advent-of-code/pkg/aoc2025/d11/p2"
	aoc2025d12p1 "github.com/zvone10/advent-of-code/pkg/aoc2025/d12/p1"
	aoc2025d12p2 "github.com/zvone10/advent-of-code/pkg/aoc2025/d12/p2"
	aoc2025d2p1 "github.com/zvone10/advent-of-code/pkg/aoc2025/d2/p1"
	aoc2025d2p2 "github.com/zvone10/advent-of-code/pkg/aoc2025/d2/p2"
	aoc2025d3p1 "github.com/zvone10/advent-of-code/pkg/aoc2025/d3/p1"
	aoc2025d3p2 "github.com/zvone10/advent-of-code/pkg/aoc2025/d3/p2"
	aoc2025d4p1 "github.com/zvone10/advent-of-code/pkg/aoc2025/d4/p1"
	aoc2025d4p2 "github.com/zvone10/advent-of-code/pkg/aoc2025/d4/p2"
	aoc2025d5p1 "github.com/zvone10/advent-of-code/pkg/aoc2025/d5/p1"
	aoc2025d5p2 "github.com/zvone10/advent-of-code/pkg/aoc2025/d5/p2"
	aoc2025d6p1 "github.com/zvone10/advent-of-code/pkg/aoc2025/d6/p1"
	aoc2025d6p2 "github.com/zvone10/advent-of-code/pkg/aoc2025/d6/p2"
	aoc2025d7p1 "github.com/zvone10/advent-of-code/pkg/aoc2025/d7/p1"
	aoc2025d7p2 "github.com/zvone10/advent-of-code/pkg/aoc2025/d7/p2"
	aoc2025d8p1 "github.com/zvone10/advent-of-code/pkg/aoc2025/d8/p1"
	aoc2025d8p2 "github.com/zvone10/advent-of-code/pkg/aoc2025/d8/p2"
	aoc2025d9p1 "github.com/zvone10/advent-of-code/pkg/aoc2025/d9/p1"
	aoc2025d9p2 "github.com/zvone10/advent-of-code/pkg/aoc2025/d9/p2"
)

var day = flag.Int("d", 0, "Day that should be executed")

func main() {
	flag.Parse()

	switch *day {
	case 1:
		internal.NewDayRunner(aoc2025d1p1.Run, aoc2025d1p2.Run).Execute()
	case 2:
		internal.NewDayRunner(aoc2025d2p1.Run, aoc2025d2p2.Run).Execute()
	case 3:
		internal.NewDayRunner(aoc2025d3p1.Run, aoc2025d3p2.Run).Execute()
	case 4:
		internal.NewDayRunner(aoc2025d4p1.Run, aoc2025d4p2.Run).Execute()
	case 5:
		internal.NewDayRunner(aoc2025d5p1.Run, aoc2025d5p2.Run).Execute()
	case 6:
		internal.NewDayRunner(aoc2025d6p1.Run, aoc2025d6p2.Run).Execute()
	case 7:
		internal.NewDayRunner(aoc2025d7p1.Run, aoc2025d7p2.Run).Execute()
	case 8:
		internal.NewDayRunner(aoc2025d8p1.Run, aoc2025d8p2.Run).Execute()
	case 9:
		internal.NewDayRunner(aoc2025d9p1.Run, aoc2025d9p2.Run).Execute()
	case 10:
		internal.NewDayRunner(aoc2025d10p1.Run, aoc2025d10p2.Run).Execute()
	case 11:
		internal.NewDayRunner(aoc2025d11p1.Run, aoc2025d11p2.Run).Execute()
	case 12:
		internal.NewDayRunner(aoc2025d12p1.Run, aoc2025d12p2.Run).Execute()
	default:
		fmt.Println("no case executed!")
	}
}
