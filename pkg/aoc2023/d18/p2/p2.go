package aoc2023d18p2

import (
	"os"
	"strconv"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY = 18
const PART = 2

var xPoints []int64
var yPoints []int64

func Run() {
	data, err := os.ReadFile("./inputs/2023/d18.txt")
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(data), "\n")
	var x, y int64
	xPoints = append(xPoints, x)
	yPoints = append(yPoints, y)
	var boundaries int64
	for _, line := range inputs {
		movement := strings.Split(line, " ")
		l, _ := strconv.ParseInt(strings.TrimSpace(movement[2])[2:7], 16, 64)

		direction := strings.TrimSpace(movement[2])[7]
		switch direction {
		case '0':
			x = x + l
		case '2':
			x = x - l
		case '3':
			y = y + l
		case '1':
			y = y - l
		}
		boundaries += l
		xPoints = append(xPoints, x)
		yPoints = append(yPoints, y)
	}

	var area int64
	numOfPoints := len(xPoints)
	for i := 0; i < numOfPoints; i++ {
		area += (xPoints[i]*yPoints[(i+1)%numOfPoints] - yPoints[i]*xPoints[(i+1)%numOfPoints])
	}
	area /= 2
	interior := abs(area) - boundaries/2 + 1
	internal.PrintResult(DAY, PART, interior+boundaries)
}

func abs(x int64) int64 {
	if x >= 0 {
		return x
	}

	return -x
}
