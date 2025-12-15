package aoc2025d9p1

import (
	"os"
	"strconv"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
	"github.com/zvone10/advent-of-code/internal/math"
)

const DAY, PART = 9, 1

func Run() {
	data, err := os.ReadFile("./inputs/2025/d9.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	coordsList := make([]coords, 0, len(lines))
	for _, line := range lines {
		n := strings.Split(line, ",")
		x, _ := strconv.Atoi(n[0])
		y, _ := strconv.Atoi(n[1])
		coordsList = append(coordsList, coords{x: x, y: y})
	}

	maxArea := 0
	for i := 0; i < len(coordsList); i++ {
		for j := i + 1; j < len(coordsList); j++ {
			c1 := coordsList[i]
			c2 := coordsList[j]

			a := (math.Abs(c1.x-c2.x) + 1) * (math.Abs(c1.y-c2.y) + 1)

			if a > maxArea {
				maxArea = a
			}
		}
	}

	internal.PrintResult(DAY, PART, maxArea)

}

type coords struct {
	x, y int
}
