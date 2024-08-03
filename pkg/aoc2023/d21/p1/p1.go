package aoc2023d21p1

import (
	"os"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY = 21
const PART = 1

var garden []string
var reachable [][]bool
var visited [][]bool
var queue []point
var rs, cs int

func Run() {
	data, err := os.ReadFile("./inputs/2023/d21.txt")
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(data), "\n")
	garden = append(garden, inputs...)

	for i, row := range garden {
		reachableRow := make([]bool, 0)
		visitedRow := make([]bool, 0)
		for j, c := range row {
			reachableRow = append(reachableRow, false)
			visitedRow = append(visitedRow, false)
			if c == 'S' {
				rs = i
				cs = j
			}
		}
		reachable = append(reachable, reachableRow)
		visited = append(visited, visitedRow)
	}

	queue = append(queue, point{r: rs, c: cs, steps: 64})
	travel()

	counter := 0
	for _, i := range reachable {
		for _, j := range i {
			if j {
				counter++
			}
		}
	}

	internal.PrintResult(DAY, PART, counter)
}

func travel() {
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if p.r < 0 || p.r >= len(garden) || p.c < 0 || p.c >= len(garden[0]) {
			continue
		}

		if garden[p.r][p.c] == '#' {
			continue
		}

		if visited[p.r][p.c] {
			continue
		}

		if p.steps%2 == 0 {
			reachable[p.r][p.c] = true
		}

		if p.steps == 0 {
			continue
		}

		visited[p.r][p.c] = true
		for _, diff := range internal.GenerateMoveDiffs() {
			rn, cn := p.r+diff.Y, p.c+diff.X
			queue = append(queue, point{r: rn, c: cn, steps: p.steps - 1})
		}
	}

}

type point struct {
	r     int
	c     int
	steps int
}
