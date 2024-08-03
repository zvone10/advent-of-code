package aoc2023d18p1

import (
	"os"
	"strconv"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY = 18
const PART = 1

var xPoints []int
var yPoints []int

func Run() {
	data, err := os.ReadFile("./inputs/2023/d18.txt")
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(data), "\n")
	x, y := 0, 0
	xPoints = append(xPoints, x)
	yPoints = append(yPoints, y)
	boundaries := 0
	for _, line := range inputs {
		movement := strings.Split(line, " ")
		l, _ := strconv.Atoi(strings.TrimSpace(movement[1]))
		switch movement[0] {
		case "R":
			x = x + l
		case "L":
			x = x - l
		case "U":
			y = y + l
		case "D":
			y = y - l
		}
		boundaries += l
		xPoints = append(xPoints, x)
		yPoints = append(yPoints, y)
	}
	area := 0
	numOfPoints := len(xPoints)
	for i := 0; i < numOfPoints; i++ {
		area += (xPoints[i]*yPoints[(i+1)%numOfPoints] - yPoints[i]*xPoints[(i+1)%numOfPoints])
	}
	area /= 2
	interior := abs(area) - boundaries/2 + 1
	internal.PrintResult(DAY, PART, interior+boundaries)
}

func abs(x int) int {
	if x >= 0 {
		return x
	}

	return -x
}
