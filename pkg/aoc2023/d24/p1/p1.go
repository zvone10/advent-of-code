package aoc2023d24p1

import (
	"os"
	"strconv"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY = 24
const PART = 1

const L, U = 200000000000000, 400000000000000

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

	count := 0
	for i, h := range hailstones {
		for j := i + 1; j < len(hailstones); j++ {
			h1, h2 := h, hailstones[j]
			A1 := h1.yv / h1.xv
			A2 := h2.yv / h2.xv

			if A1 == A2 {
				continue
			}

			x := (A1*h1.x - A2*h2.x + h2.y - h1.y) / (A1 - A2)
			y := A1*(x-h1.x) + h1.y

			tx1 := (x - h1.x) / h.xv
			ty1 := (y - h1.y) / h.yv

			tx2 := (x - h2.x) / h2.xv
			ty2 := (y - h2.y) / h2.yv

			if x >= L && x <= U && y >= L && y <= U && tx1 > 0 && ty1 > 0 && tx2 > 0 && ty2 > 0 {
				count++
			}
		}
	}

	internal.PrintResult(DAY, PART, count)
}

type hailstone struct {
	x, y, z, xv, yv, zv float64
}
