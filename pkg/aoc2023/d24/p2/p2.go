package aoc2023d24p2

import (
	"os"
	"strconv"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY = 24
const PART = 2

const L, U = 7, 27

var hailstones []hailstone

func Run() {
	data, err := os.ReadFile("./inputs/2023/d24.txt")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(data), "\n")
	for _, r := range rows {
		t := strings.Split(r, " @ ")
		coords := strings.Split(t[0], ", ")
		speeds := strings.Split(t[1], ", ")
		x, _ := strconv.ParseFloat(strings.TrimPrefix(coords[0], " "), 64)
		y, _ := strconv.ParseFloat(strings.TrimPrefix(coords[1], " "), 64)
		z, _ := strconv.ParseFloat(strings.TrimPrefix(coords[2], " "), 64)

		xv, _ := strconv.ParseFloat(strings.TrimPrefix(speeds[0], " "), 64)
		yv, _ := strconv.ParseFloat(strings.TrimPrefix(speeds[1], " "), 64)
		zv, _ := strconv.ParseFloat(strings.TrimPrefix(speeds[2], " "), 64)

		hailstones = append(hailstones, hailstone{
			x:  x,
			y:  y,
			z:  z,
			xv: xv,
			yv: yv,
			zv: zv,
		})
	}

	var x, y, z float64

	internal.PrintResult(DAY, PART, x+y+z)
}

type hailstone struct {
	x, y, z, xv, yv, zv float64
}
