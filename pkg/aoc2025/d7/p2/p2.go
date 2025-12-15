package aoc2025d7p2

import (
	"os"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY, PART = 7, 2

var dr, dc int

var manifold []string
var cache map[point]int

func Run() {
	data, err := os.ReadFile("./inputs/2025/d7.txt")
	if err != nil {
		panic(err)
	}

	manifold = strings.Split(string(data), "\n")

	dr, dc = len(manifold), len(manifold[0])
	cache = make(map[point]int)

	var sr, sc int
	for c := 0; c < len(manifold[0]); c++ {
		if manifold[0][c] == 'S' {
			sr, sc = 0, c
			break
		}
	}

	internal.PrintResult(DAY, PART, search(sr, sc))
}

type point struct {
	r, c int
}

func search(r, c int) int {
	value, ok := cache[point{r, c}]
	if ok {
		return value
	}

	if r >= 0 && r < dr && c >= 0 && c < dc {
		if manifold[r][c] == 'S' || manifold[r][c] == '.' {
			return search(r+1, c)
		} else if manifold[r][c] == '^' {
			v := search(r, c-1) + search(r, c+1)
			cache[point{r: r, c: c}] = v
			return v
		}
	}

	return 1
}
