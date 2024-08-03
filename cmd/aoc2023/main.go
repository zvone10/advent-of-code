package main

import (
	"flag"
	"fmt"

	"github.com/zvone10/advent-of-code/internal"
	aoc2023d1p1 "github.com/zvone10/advent-of-code/pkg/aoc2023/d1/p1"
	aoc2023d1p2 "github.com/zvone10/advent-of-code/pkg/aoc2023/d1/p2"
	aoc2023d18p1 "github.com/zvone10/advent-of-code/pkg/aoc2023/d18/p1"
	aoc2023d18p2 "github.com/zvone10/advent-of-code/pkg/aoc2023/d18/p2"
	aoc2023d19p1 "github.com/zvone10/advent-of-code/pkg/aoc2023/d19/p1"
	aoc2023d19p2 "github.com/zvone10/advent-of-code/pkg/aoc2023/d19/p2"
	aoc2023d2p1 "github.com/zvone10/advent-of-code/pkg/aoc2023/d2/p1"
	aoc2023d2p2 "github.com/zvone10/advent-of-code/pkg/aoc2023/d2/p2"
	aoc2023d20p1 "github.com/zvone10/advent-of-code/pkg/aoc2023/d20/p1"
	aoc2023d20p2 "github.com/zvone10/advent-of-code/pkg/aoc2023/d20/p2"
	aoc2023d21p1 "github.com/zvone10/advent-of-code/pkg/aoc2023/d21/p1"
	aoc2023d21p2 "github.com/zvone10/advent-of-code/pkg/aoc2023/d21/p2"
	aoc2023d22p1 "github.com/zvone10/advent-of-code/pkg/aoc2023/d22/p1"
	aoc2023d22p2 "github.com/zvone10/advent-of-code/pkg/aoc2023/d22/p2"
	aoc2023d23p1 "github.com/zvone10/advent-of-code/pkg/aoc2023/d23/p1"
	aoc2023d23p2 "github.com/zvone10/advent-of-code/pkg/aoc2023/d23/p2"
	aoc2023d24p1 "github.com/zvone10/advent-of-code/pkg/aoc2023/d24/p1"
	aoc2023d24p2 "github.com/zvone10/advent-of-code/pkg/aoc2023/d24/p2"
)

var day = flag.Int("d", 0, "Day that should be executed")

func main() {
	flag.Parse()

	switch *day {
	case 1:
		internal.NewDayRunner(aoc2023d1p1.Run, aoc2023d1p2.Run).Execute()
	case 2:
		internal.NewDayRunner(aoc2023d2p1.Run, aoc2023d2p2.Run).Execute()
	case 3:
		fmt.Println("empty")
	case 18:
		internal.NewDayRunner(aoc2023d18p1.Run, aoc2023d18p2.Run).Execute()
	case 19:
		internal.NewDayRunner(aoc2023d19p1.Run, aoc2023d19p2.Run).Execute()
	case 20:
		internal.NewDayRunner(aoc2023d20p1.Run, aoc2023d20p2.Run).Execute()
	case 21:
		internal.NewDayRunner(aoc2023d21p1.Run, aoc2023d21p2.Run).Execute()
	case 22:
		internal.NewDayRunner(aoc2023d22p1.Run, aoc2023d22p2.Run).Execute()
	case 23:
		internal.NewDayRunner(aoc2023d23p1.Run, aoc2023d23p2.Run).Execute()
	case 24:
		internal.NewDayRunner(aoc2023d24p1.Run, aoc2023d24p2.Run).Execute()
	default:
		fmt.Println("no case executed!")
	}
}
