package aoc2023d23p1

import (
	"os"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY = 23
const PART = 1

var graph []string
var visited map[int]bool
var longestWalk int
var moves []internal.Coordinates

func Run() {
	data, err := os.ReadFile("./inputs/2023/d23.txt")
	if err != nil {
		panic(err)
	}

	graph = strings.Split(string(data), "\n")
	//fmt.Println(inputs)

	startC := 0

	for i := 0; i < len(graph[0]); i++ {
		if graph[0][i] == '.' {
			startC = i
			break
		}
	}

	moves = internal.GenerateMoveDiffs()

	visited = make(map[int]bool)
	travel(startC, 0, 0)

	internal.PrintResult(DAY, PART, longestWalk)
}

func travel(x, y, length int) {
	if x < 0 || x >= len(graph[0]) || y < 0 || y >= len(graph) {
		return
	}

	if graph[y][x] == '#' {
		return
	}

	if v, ok := visited[y*len(graph[0])+x]; ok && v {
		return
	}

	if y == len(graph)-1 && graph[y][x] == '.' && length > longestWalk {
		longestWalk = length
		return
	}

	if y != len(graph)-1 {
		visited[y*len(graph[0])+x] = true
	}

	if graph[y][x] == '.' {
		for _, c := range moves {
			travel(x+c.X, y+c.Y, length+1)
		}
	} else if graph[y][x] == '>' {
		travel(x+1, y, length+1)
	} else if graph[y][x] == '<' {
		travel(x-1, y, length+1)
	} else if graph[y][x] == 'v' {
		travel(x, y+1, length+1)
	} else if graph[y][x] == '^' {
		travel(x, y-1, length+1)
	}

	visited[y*len(graph[0])+x] = false
}

func copy(source map[int]bool) map[int]bool {
	c := make(map[int]bool)
	for k, v := range source {
		c[k] = v
	}
	return c
}
