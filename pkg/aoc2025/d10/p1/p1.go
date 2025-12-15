package aoc2025d10p1

import (
	"os"
	"strconv"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY, PART = 10, 1

var buttons [][]int

func Run() {
	data, err := os.ReadFile("./inputs/2025/d10.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	//fmt.Println(lines)

	sum := 0
	for _, line := range lines {
		tokens := strings.Split(line, " ")

		target := tokens[0]
		buttons = make([][]int, 0)
		for i := 1; i < len(tokens); i++ {
			if tokens[i][0] == '(' {
				b := make([]int, 0)
				text := tokens[i][1 : len(tokens[i])-1]
				numStrings := strings.Split(text, ",")
				for _, ns := range numStrings {
					n, _ := strconv.Atoi(ns)
					b = append(b, n)
				}
				buttons = append(buttons, b)
			}
		}

		var targetNumber int
		// convert #..##### to int where .->0 and #->1
		for i := 1; i < len(target)-1; i++ {
			if target[i] == '#' {
				targetNumber += 1 << (i - 1)
			}
		}
		sum += find(targetNumber)
	}
	internal.PrintResult(DAY, PART, sum)

}

func find(target int) int {
	processed := make(map[int]int, 0)

	queue := make([]node, 1)
	queue = append(queue, node{
		state: 0,
		dist:  0,
	})

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		state := current.state

		if v, ok := processed[state]; ok {
			if v < current.dist {
				continue
			}
		}

		if state == target {
			return current.dist
		}

		processed[state] = current.dist

		for _, b := range buttons {
			newState := state
			for _, n := range b {
				newState = flipBit(newState, n)
			}
			queue = append(queue, node{
				state: newState,
				dist:  current.dist + 1,
			})
		}
	}
	return -1
}

type node struct {
	state int
	dist  int
}

func flipBit(x int, bit int) int {
	return x ^ (1 << bit)
}
