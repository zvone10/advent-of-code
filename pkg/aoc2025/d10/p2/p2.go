package aoc2025d10p2

import (
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY, PART = 10, 2

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

		var targetState []int
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
			} else if tokens[i][0] == '{' {
				text := tokens[i][1 : len(tokens[i])-1]

				numStrings := strings.Split(text, ",")
				targetState = make([]int, 0, len(numStrings))
				for _, ns := range numStrings {
					n, _ := strconv.Atoi(ns)
					targetState = append(targetState, n)
				}
			}

		}
		sum += find(targetState)
	}
	internal.PrintResult(DAY, PART, sum)

}

func find(target []int) int {
	processed := make(map[string]int, 0)

	queue := make([]node, 0)
	queue = append(queue, node{
		state: make([]int, len(target)),
		dist:  0,
	})

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		state := current.state
		stateKey := sliceToIntKey(state)

		if v, ok := processed[stateKey]; ok {
			if v < current.dist {
				continue
			}
		}

		if slices.Equal(state, target) {
			return current.dist
		}

		processed[stateKey] = current.dist

		for _, b := range buttons {
			newState := make([]int, len(state))
			copy(newState, state)
			for _, n := range b {
				newState[n] = newState[n] + 1
			}
			queue = append(queue, node{
				state: newState,
				dist:  current.dist + 1,
			})
			//fmt.Println(queue)
		}
	}
	return -1
}

type node struct {
	state []int
	dist  int
}

func sliceToIntKey(s []int) string {
	str := make([]string, len(s))
	for i, v := range s {
		// Convert the int to a string
		str[i] = strconv.Itoa(v)
	}
	// Join all parts with a unique separator, like "|" or ","
	return strings.Join(str, "|")
}
