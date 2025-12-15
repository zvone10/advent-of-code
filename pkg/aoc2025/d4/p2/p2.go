package aoc2025d4p2

import (
	"fmt"
	"os"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY, PART = 4, 2

var layout []string
var nr, nc int

func Run() {
	data, err := os.ReadFile("./inputs/2025/d4.txt")
	if err != nil {
		panic(err)
	}

	layout = strings.Split(string(data), "\n")
	nr, nc = len(layout), len(layout[0])

	numOfSpots := 0
	shouldContinue := true
	for shouldContinue {
		newLayout := copyLayout()
		shouldContinue = false
		for i, r := range layout {
			for j, c := range r {
				if c == '@' {
					counter := 0
					for ni := -1; ni <= 1; ni++ {
						for nj := -1; nj <= 1; nj++ {
							if !(ni == 0 && nj == 0) && isInside(i+ni, j+nj) {
								if layout[i+ni][j+nj] == '@' {
									counter++
								}
							}
						}
					}
					if counter < 4 {
						newLayout[i] = newLayout[i][:j] + "." + newLayout[i][j+1:]
						numOfSpots++
						shouldContinue = true
					}
				}
			}
		}
		fmt.Println("Iteration complete")
		//fmt.Println(strings.Join(newLayout, "\n"))
		layout = newLayout
	}

	internal.PrintResult(DAY, PART, numOfSpots)
}

func copyLayout() []string {
	newLayout := make([]string, nr)
	for i := 0; i < nr; i++ {
		newLayout[i] = layout[i]
	}
	return newLayout
}

func isInside(i, j int) bool {
	if i >= 0 && i < nr && j >= 0 && j < nc {
		return true
	}
	return false
}
