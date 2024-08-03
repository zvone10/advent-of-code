package d17p1

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type node struct {
	r                int
	c                int
	dr               int
	dc               int
	consecutiveSteps int
}

var graph [][]int
var q []node
var costs map[node]int
var visited map[node]bool

type direction struct {
	dr int
	dc int
}

var directions []direction

func Run() {
	data, err := os.ReadFile("./inputs/d17.txt")
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(data), "\n")
	for _, input := range inputs {
		var row []int

		for _, char := range input {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Println(err)
				return
			}

			row = append(row, num)
		}
		graph = append(graph, row)
	}

	//consecutives = make(map[node]int)
	costs = make(map[node]int)

	directions = []direction{
		direction{dr: 0, dc: 1},
		direction{dr: 0, dc: -1},
		direction{dr: -1, dc: 0},
		direction{dr: 1, dc: 0},
	}

	n := node{
		r:                0,
		c:                0,
		dr:               0,
		dc:               0,
		consecutiveSteps: 0,
	}
	q = append(q, n)
	costs[n] = 0
	visited = map[node]bool{}

	for len(q) > 0 {
		index, n := findMinQueue()
		cost := costs[n]
		q = append(q[0:index], q[index+1:]...)

		_, v := visited[n]
		if v {
			continue
		}
		visited[n] = true

		for _, d := range directions {
			r := n.r + d.dr
			c := n.c + d.dc

			if r < 0 || r >= len(graph) || c < 0 || c >= len(graph[0]) {
				continue
			}

			if (d.dr == -n.dr && d.dr != 0) || (d.dc == -n.dc && d.dc != 0) {
				continue
			}

			consecutives := 0
			if d.dr == n.dr && d.dc == n.dc {
				consecutives = n.consecutiveSteps + 1
			}

			if consecutives == 3 {
				continue
			}

			newNode := node{
				r:                r,
				c:                c,
				dr:               d.dr,
				dc:               d.dc,
				consecutiveSteps: consecutives,
			}
			_, newNodeCostExists := costs[newNode]

			if !newNodeCostExists {
				costs[newNode] = cost + graph[r][c]
				q = append(q, newNode)
			}
		}

	}

	minimum := math.MaxInt
	for n, c := range costs {
		if n.c == len(graph[0])-1 && n.r == len(graph)-1 {
			fmt.Println(n)
			fmt.Println(c)

			if c < minimum {
				minimum = c
			}
		}
	}
	fmt.Println(minimum)
}

func findMinQueue() (int, node) {
	cost := -1
	var minNode node
	var minIndex int
	for i, node := range q {
		c := costs[node]
		if cost == -1 || c < cost {
			cost = c
			minNode = node
			minIndex = i
		}
	}
	return minIndex, minNode
}
