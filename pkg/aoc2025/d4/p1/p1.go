package aoc2025d4p1

import (
	"os"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY, PART = 4, 1

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
					numOfSpots++
				}

			}
		}
	}
	internal.PrintResult(DAY, PART, numOfSpots)

}

func isInside(i, j int) bool {
	if i >= 0 && i < nr && j >= 0 && j < nc {
		return true
	}
	return false
}
