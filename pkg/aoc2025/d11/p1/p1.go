package aoc2025d11p1

import (
	"os"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY, PART = 11, 1

var graph map[string][]string

var count int

func Run() {
	data, err := os.ReadFile("./inputs/2025/d11.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	graph = make(map[string][]string)

	for _, l := range lines {
		tokens := strings.Split(l, ": ")
		//fmt.Println(tokens)
		key := tokens[0]
		values := strings.Split(tokens[1], " ")
		graph[key] = values
	}

	traverse("you")

	internal.PrintResult(DAY, PART, count)
}

func traverse(node string) {
	if node == "out" {
		count++
		return
	}

	for _, child := range graph[node] {
		traverse(child)
	}
}
