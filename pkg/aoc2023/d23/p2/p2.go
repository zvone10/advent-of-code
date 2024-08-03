package aoc2023d23p2

import (
	"fmt"
	"os"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY = 23
const PART = 2

var graph []string
var points map[int]bool
var compressedGraph map[int][]node
var visited map[int]bool
var longestWalk int
var moves []internal.Coordinates
var endNode int

func Run() {
	data, err := os.ReadFile("./inputs/2023/d23.txt")
	if err != nil {
		panic(err)
	}

	graph = strings.Split(string(data), "\n")

	startC := 0
	endC := 0

	for i := 0; i < len(graph[0]); i++ {
		if graph[0][i] == '.' {
			startC = i
		}

		if graph[len(graph)-1][i] == '.' {
			endC = i
		}
	}

	endNode = (len(graph)-1)*(len(graph[0])) + endC

	moves = internal.GenerateMoveDiffs()

	points = make(map[int]bool)
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[0]); j++ {
			if graph[i][j] != '#' {
				surrounding := 0
				for _, m := range moves {
					x := j + m.X
					y := i + m.Y
					if x >= 0 && x < len(graph[0]) && y >= 0 && y < len(graph) && graph[y][x] != '#' {
						surrounding++
					}
				}

				if surrounding >= 3 {
					points[i*len(graph[0])+j] = true
				}
			}
		}
	}

	points[startC] = true
	points[len(graph[0])*(len(graph)-1)+endC] = true

	fmt.Println(points)

	compressedGraph = make(map[int][]node)

	for p, _ := range points {
		visited = make(map[int]bool)
		flood(p%len(graph[0]), p/len(graph[0]), p%len(graph[0]), p/len(graph[0]), 0)
	}
	fmt.Println(compressedGraph)

	visited = make(map[int]bool)
	travel(startC, 0)

	internal.PrintResult(DAY, PART, longestWalk)
}

func flood(xstart, ystart, x, y, length int) {
	// if xstart == 1 && ystart == 0 {
	// 	fmt.Println("$$$")
	// 	fmt.Println(x, y)
	// }

	if x < 0 || x >= len(graph[0]) || y < 0 || y >= len(graph) {
		return
	}

	if graph[y][x] == '#' {
		return
	}

	if _, ok := visited[y*len(graph[0])+x]; ok {
		return
	}

	if _, ok := points[y*len(graph[0])+x]; ok && length != 0 {
		compressedGraph[ystart*len(graph[0])+xstart] = append(compressedGraph[ystart*len(graph[0])+xstart], node{
			point:    y*len(graph[0]) + x,
			distance: length,
		})
		return
	}

	visited[y*len(graph[0])+x] = true
	for _, c := range moves {
		flood(xstart, ystart, x+c.X, y+c.Y, length+1)
	}
}

func travel(node, length int) {
	if v, ok := visited[node]; ok && v {
		return
	}

	if node == endNode {
		//fmt.Println(length)
		if length > longestWalk {
			longestWalk = length
		}
		return
	}

	if node != endNode {
		visited[node] = true
	}

	for _, n := range compressedGraph[node] {
		travel(n.point, length+n.distance)
	}

	visited[node] = false
}

type node struct {
	point    int
	distance int
}
