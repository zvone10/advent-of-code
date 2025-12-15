package aoc2025d11p2

import (
	"os"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY, PART = 11, 2

var graph map[string][]string
var cache map[string]int

func Run() {
	data, err := os.ReadFile("./inputs/2025/d11.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	graph = make(map[string][]string)
	cache = make(map[string]int)

	for _, l := range lines {
		tokens := strings.Split(l, ": ")
		//fmt.Println(tokens)
		key := tokens[0]
		values := strings.Split(tokens[1], " ")
		graph[key] = values
	}

	svrFft := traverse("svr", "fft")
	fftDac := traverse("fft", "dac")
	dacOut := traverse("dac", "out")

	svrDac := traverse("svr", "dac")
	dacFft := traverse("dac", "fft")
	fftOut := traverse("fft", "out")

	internal.PrintResult(DAY, PART, svrFft*fftDac*dacOut+svrDac*dacFft*fftOut)
}

func traverse(node string, dest string) int {
	if v, ok := cache[node+dest]; ok {
		return v
	}

	if node == dest {
		return 1
	}

	sum := 0
	for _, child := range graph[node] {
		sum += traverse(child, dest)
	}
	cache[node+dest] = sum
	return sum
}
