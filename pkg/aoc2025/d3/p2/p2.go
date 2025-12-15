package aoc2025d3p2

import (
	"os"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY, PART = 3, 2

var cache map[key]uint64

type key struct {
	s     string
	depth int
}

func Run() {
	data, err := os.ReadFile("./inputs/2025/d3.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	cache = make(map[key]uint64)

	sum := uint64(0)
	for _, l := range lines {
		v := optimalNumber(l, 12)
		sum += v
	}

	internal.PrintResult(DAY, PART, sum)
}

func optimalNumber(a string, depth int) uint64 {
	if depth == 0 || len(a) == 0 {
		return uint64(0)
	}

	if v, ok := cache[key{a, depth}]; ok {
		return v
	}

	if len(a) == depth {
		v := uint64(0)
		for _, c := range a {
			v = 10*v + uint64(c-'0')
		}
		cache[key{a, depth}] = v
		return v
	}

	maxValue := uint64(0)
	current := uint64(a[0] - '0')
	p := pow10(depth - 1)
	for i := 1; len(a)-i >= depth; i++ {
		v := p*current + optimalNumber(a[i:], depth-1)
		if v > maxValue {
			maxValue = v
		}
	}

	v := optimalNumber(a[1:], depth)
	if v > maxValue {
		maxValue = v
	}

	cache[key{a, depth}] = maxValue
	return maxValue
}

func pow10(n int) uint64 {
	p := uint64(1)
	for i := 0; i < n; i++ {
		p = p * 10
	}
	return p
}
