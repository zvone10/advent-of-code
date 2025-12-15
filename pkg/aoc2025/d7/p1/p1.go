package aoc2025d7p1

import (
	"fmt"
	"os"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY, PART = 7, 1

var manifold []string

func Run() {
	data, err := os.ReadFile("./inputs/2025/d7.txt")
	if err != nil {
		panic(err)
	}

	manifold = strings.Split(string(data), "\n")

	dr, dc := len(manifold), len(manifold[0])

	var sr, sc int
	for i := 0; i < len(manifold); i++ {
		for c := 0; c < len(manifold[i]); c++ {
			if manifold[i][c] == 'S' {
				sr, sc = i, c
			}
		}
	}

	queue := make([]point, 0, 1)
	queue = append(queue, point{r: sr, c: sc})
	filled := make(map[point]bool)

	numOfSplits := 0
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		ok, isFilled := filled[p]
		if ok && isFilled {
			continue
		}
		if p.r >= 0 && p.r < dr && p.c >= 0 && p.c < dc && !filled[p] {
			if manifold[p.r][p.c] == 'S' || manifold[p.r][p.c] == '.' {
				queue = append(queue, point{
					r: p.r + 1,
					c: p.c,
				})
			} else if manifold[p.r][p.c] == '^' {
				queue = append(queue, point{
					r: p.r,
					c: p.c - 1,
				})

				queue = append(queue, point{
					r: p.r,
					c: p.c + 1,
				})
				numOfSplits++
				fmt.Println("split at:", p, " total splits:", numOfSplits)
			}
		}
		filled[p] = true
	}

	internal.PrintResult(DAY, PART, numOfSplits)
}

type point struct {
	r, c int
}
