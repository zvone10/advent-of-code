package aoc2023d22p1

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY = 22
const PART = 2

var bricks []brick
var graph map[int][]int

func Run() {
	data, err := os.ReadFile("./inputs/2023/d22.txt")
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(data), "\n")

	for _, i := range inputs {
		edges := strings.Split(i, "~")
		b1Coords := *internal.StringToIntSlice(edges[0], ",")
		b2Coords := *internal.StringToIntSlice(edges[1], ",")

		c1 := cube{x: b1Coords[0], y: b1Coords[1], z: b1Coords[2]}
		c2 := cube{x: b2Coords[0], y: b2Coords[1], z: b2Coords[2]}

		var cBottom, cTop cube
		if c1.z <= c2.z {
			cBottom = c1
			cTop = c2
		} else {
			cBottom = c2
			cTop = c1
		}

		bricks = append(bricks, brick{
			c1: cBottom,
			c2: cTop,
		})
	}

	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].c1.z < bricks[j].c1.z
	})

	for i, bStart := range bricks {
		maxZ := 0
		for j := i - 1; j >= 0; j-- {
			bOther := bricks[j]

			if intersect(bOther.c1.y, bOther.c2.y, bOther.c1.x, bOther.c2.x, bStart.c1.y, bStart.c2.y, bStart.c1.x, bStart.c2.x) {
				if bOther.c2.z+1 > maxZ {
					maxZ = bOther.c2.z + 1
				}
			}

		}

		bricks[i].c2.z = bStart.c2.z + maxZ - bStart.c1.z
		bricks[i].c1.z = maxZ
	}

	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].c1.z < bricks[j].c1.z
	})

	graph = make(map[int][]int)
	gSupport := make(map[int][]int)

	for i, b := range bricks {
		graph[i] = []int{}
		for j := i + 1; j < len(bricks); j++ {
			bOther := bricks[j]
			if bOther.c1.z == b.c2.z+1 && intersect(bOther.c1.y, bOther.c2.y, bOther.c1.x, bOther.c2.x, b.c1.y, b.c2.y, b.c1.x, b.c2.x) {
				graph[i] = append(graph[i], j)
				gSupport[j] = append(gSupport[j], i)
			}
		}
	}

	fmt.Println(graph)
	result := 0
	for i := range bricks {
		q := []int{i}

		removed := make(map[int]bool)
		removed[i] = true

		for len(q) > 0 {
			h := q[0]
			q = q[1:]

			for _, b := range graph[h] {
				supportedBy := gSupport[b]
				numOfRemaining := len(supportedBy)

				for _, x := range supportedBy {
					if _, ok := removed[x]; ok {
						numOfRemaining--
					}
				}

				if numOfRemaining == 0 {
					removed[b] = true
					q = append(q, b)
					//fmt.Println("?????")
					//fmt.Println(q)
				}
			}

		}

		fmt.Println(removed)
		result += (len(removed) - 1)
	}

	// for _, v := range graph {
	// 	canBeRemoved := true
	// 	for _, b := range v {
	// 		if len(gSupport[b]) < 2 {
	// 			canBeRemoved = false
	// 			break
	// 		}
	// 	}
	// 	if canBeRemoved {
	// 		result++
	// 	}
	// }
	internal.PrintResult(DAY, PART, result)
}

type brick struct {
	c1 cube
	c2 cube
}

type cube struct {
	x, y, z int
}

type Rectangle struct {
	Left, Top, Right, Bottom int
}

func contains(x int, arr *[]int) bool {
	for _, a := range *arr {
		if x == a {
			return true
		}
	}
	return false
}

func intersect(r1Left, r1Right, r1Bottom, r1Top, r2Left, r2Right, r2Bottom, r2Top int) bool {
	// Check horizontal overlap
	//fmt.Println(r1Left, r1Right, r1Bottom, r1Top, r2Left, r2Right, r2Bottom, r2Top)
	if min(r1Right, r2Right) < max(r1Left, r2Left) {
		return false
	}
	// Check vertical overlap
	if min(r1Top, r2Top) < max(r1Bottom, r2Bottom) {
		return false
	}

	//fmt.Println(true)
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
