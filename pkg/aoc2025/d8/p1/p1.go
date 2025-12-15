package aoc2025d8p1

import (
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY, PART = 8, 1

func Run() {
	data, err := os.ReadFile("./inputs/2025/d8.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	boxes := make([]box, 0, len(lines))
	for _, line := range lines {
		n := strings.Split(line, ",")
		x, _ := strconv.Atoi(n[0])
		y, _ := strconv.Atoi(n[1])
		z, _ := strconv.Atoi(n[2])
		boxes = append(boxes, box{x: x, y: y, z: z})
	}

	pairs := make([]pair, 0)
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			b1 := boxes[i]
			b2 := boxes[j]
			distance := square(b1.x-b2.x) + square(b1.y-b2.y) + square(b1.z-b2.z)
			pairs = append(pairs, pair{
				b1:       b1,
				b2:       b2,
				distance: distance,
			})
		}
	}

	sort.Slice(pairs, func(i int, j int) bool {
		return pairs[i].distance < pairs[j].distance
	})

	group := make(map[box]int)
	g := 0
	for i := 0; i < 1000; i++ {
		p := pairs[i]

		v1, ok1 := group[p.b1]
		v2, ok2 := group[p.b2]

		if ok1 && ok2 && v1 == v2 {
			continue
		} else if ok1 && ok2 && v1 != v2 {
			for k, v := range group {
				if v == v2 {
					group[k] = v1
				}
			}
		} else if ok1 && !ok2 {
			group[p.b2] = v1
		} else if ok2 && !ok1 {
			group[p.b1] = v2
		} else {
			group[p.b1] = g
			group[p.b2] = g
			g++
		}
	}

	groupCount := make(map[int]int)
	for _, v := range group {
		groupCount[v]++
	}

	counts := make([]int, 0, len(groupCount))
	for _, v := range groupCount {
		counts = append(counts, v)
	}

	sort.Ints(counts)

	l := len(counts)
	internal.PrintResult(DAY, PART, counts[l-1]*counts[l-2]*counts[l-3])

}

type box struct {
	x, y, z int
}

type pair struct {
	b1, b2   box
	distance int
}

func square(n int) int {
	return n * n
}
