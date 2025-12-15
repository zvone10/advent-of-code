package aoc2025d1p2

import (
	"os"
	"strconv"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY, PART = 1, 2

func Run() {
	data, err := os.ReadFile("./inputs/2025/d1.txt")
	if err != nil {
		panic(err)
	}

	position := 50
	count := 0
	moves := strings.Split(string(data), "\n")
	for _, move := range moves {
		d, _ := strconv.Atoi(move[1:])

		//p100 := position / 100

		if move[0] == 'L' {
			for i := position - 1; i > position-d; i-- {
				if i%100 == 0 {
					count++
				}
			}
			position -= d
		} else {
			for i := position + 1; i < position+d; i++ {
				if i%100 == 0 {
					count++
				}
			}
			position += d
		}

		if position%100 == 0 {
			count++
		}
	}

	internal.PrintResult(DAY, PART, count)
}
