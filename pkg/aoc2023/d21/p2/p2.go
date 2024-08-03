package aoc2023d21p2

import (
	"fmt"
	"os"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY = 21
const PART = 2

const steps = 26501365 // = 65+202300*131

var garden []string

func Run() {
	data, err := os.ReadFile("./inputs/2023/d21.txt")
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(data), "\n")
	garden = append(garden, inputs...)

	// travel()

	fmt.Printf("Garden dimensions rows: %d, columns %d\n", len(garden), len(garden[0]))
	fmt.Printf("Number of fully contained squares: %d, columns %d\n", steps/len(garden), steps/len(garden[0]))

	gardenSize := len(garden)               //=131
	numOfFullSquares := steps / len(garden) // 202300

	//calculating number of full squares
	squares := numOfFullSquares - 1
	numOfOddSquares := (2*(squares/2) + 1) * (2*(squares/2) + 1)
	numOfEvenSquares := 4 * ((squares + 1) / 2) * ((squares + 1) / 2)

	fmt.Println(numOfOddSquares, numOfEvenSquares)

	counter := 0
	counter += numOfEvenSquares*travel(gardenSize/2, gardenSize/2, gardenSize-1) + numOfOddSquares*travel(gardenSize/2, gardenSize/2, gardenSize)

	//tips
	counter += travel(gardenSize-1, gardenSize/2, gardenSize-1) //top
	counter += travel(gardenSize/2, 0, gardenSize-1)
	counter += travel(gardenSize/2, gardenSize-1, gardenSize-1)
	counter += travel(0, gardenSize/2, gardenSize-1)

	numOfStepsForCornerTriangle := (gardenSize-1)/2 - 1
	counter += numOfFullSquares * (travel(gardenSize-1, 0, numOfStepsForCornerTriangle) +
		travel(gardenSize-1, gardenSize-1, numOfStepsForCornerTriangle) +
		travel(0, gardenSize-1, numOfStepsForCornerTriangle) +
		travel(0, 0, numOfStepsForCornerTriangle))

	numberOfStepsForPartialSquare := (gardenSize-1)/2 - 1 + gardenSize
	counter += squares * (travel(gardenSize-1, 0, numberOfStepsForPartialSquare) +
		travel(gardenSize-1, gardenSize-1, numberOfStepsForPartialSquare) +
		travel(0, gardenSize-1, numberOfStepsForPartialSquare) +
		travel(0, 0, numberOfStepsForPartialSquare))

	//counter += gar

	internal.PrintResult(DAY, PART, counter)
}

func travel(rowStart, columnStart, steps int) int {
	var reachable [][]bool
	var visited [][]bool
	var queue []point

	for _, row := range garden {
		reachableRow := make([]bool, len(row))
		visitedRow := make([]bool, len(row))

		reachable = append(reachable, reachableRow)
		visited = append(visited, visitedRow)
	}

	queue = append(queue, point{r: rowStart, c: columnStart, steps: steps})

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

	counter := 0
	for _, i := range reachable {
		for _, j := range i {
			if j {
				counter++
			}
		}
	}

	return counter
}

type point struct {
	r     int
	c     int
	steps int
}
