package aoc2023d19p2

import (
	"os"
	"strconv"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY = 19
const PART = 2

const (
	x = 0
	m = 1
	a = 2
	s = 3
)

type rule struct {
	index    int // x,m,a,s
	isLarger bool
	value    int
	workflow string
}

var workflows map[string][]*rule
var defaultWorkflows map[string]string
var indexMap map[byte]int

//var parts [][]int

func Run() {
	data, err := os.ReadFile("./inputs/2023/d19.txt")
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(data), "\n")

	indexMap = map[byte]int{
		'x': x,
		'm': m,
		'a': a,
		's': s,
	}
	workflows = make(map[string][]*rule)
	defaultWorkflows = make(map[string]string)

	for _, input := range inputs {
		if len(input) > 0 {
			tokens := strings.Split(input, "{")
			rules := strings.Split(tokens[1][0:len(tokens[1])-1], ",")

			for i, r := range rules {
				if i < len(rules)-1 {
					index := indexMap[r[0]]
					isLarge := (r[1] == '>')
					ruleTokens := strings.Split(r[2:], ":")
					v, _ := strconv.Atoi(ruleTokens[0])

					r := &rule{
						index:    index,
						isLarger: isLarge,
						value:    v,
						workflow: ruleTokens[1],
					}
					workflows[tokens[0]] = append(workflows[tokens[0]], r)
				} else {
					defaultWorkflows[tokens[0]] = r
				}

			}
		} else {
			break
		}
	}

	initialInterval := interval{
		min: 1,
		max: 4000,
	}
	allIntervals := []interval{
		initialInterval,
		initialInterval,
		initialInterval,
		initialInterval,
	}
	combinations := countCombination(allIntervals, "in")

	internal.PrintResult(DAY, PART, combinations)

}

func countCombination(intervals []interval, ruleName string) uint64 {

	if ruleName == "R" {
		return 0
	}

	if ruleName == "A" {
		return (intervals[x].max - intervals[x].min + 1) *
			(intervals[m].max - intervals[m].min + 1) *
			(intervals[a].max - intervals[a].min + 1) *
			(intervals[s].max - intervals[s].min + 1)
	}

	var total uint64
	currentRules := workflows[ruleName]

	isEmpty := false
	for _, r := range currentRules {
		current := intervals[r.index]

		if r.isLarger {
			lowest := max(uint64(r.value)+1, current.min)
			if lowest < current.max {
				p := interval{
					min: lowest,
					max: current.max,
				}
				newIntervals := make([]interval, 4)
				for i := 0; i < 4; i++ {
					if i != r.index {
						newIntervals[i] = intervals[i]
					} else {
						newIntervals[i] = p
					}
				}
				total += countCombination(newIntervals, r.workflow)
			}

			highest := min(current.max, uint64(r.value))
			if current.min < highest {
				intervals[r.index] = interval{
					min: current.min,
					max: highest,
				}
			} else {
				isEmpty = true
				break
			}

		} else {
			highest := min(uint64(r.value)-1, current.max)
			if highest > current.min {
				p := interval{
					min: current.min,
					max: highest,
				}
				newIntervals := make([]interval, 4)
				for i := 0; i < 4; i++ {
					if i != r.index {
						newIntervals[i] = intervals[i]
					} else {
						newIntervals[i] = p
					}
				}
				total += countCombination(newIntervals, r.workflow)
			}

			lowest := max(current.min, uint64(r.value))
			if current.max > lowest {
				intervals[r.index] = interval{
					min: lowest,
					max: current.max,
				}
			} else {
				isEmpty = true
				break
			}
		}
	}

	if !isEmpty {
		total += countCombination(intervals, defaultWorkflows[ruleName])
	}

	return total
}

type interval struct {
	min uint64
	max uint64
}

func min(x, y uint64) uint64 {
	if x < y {
		return x
	}

	return y
}

func max(x, y uint64) uint64 {
	if x > y {
		return x
	}

	return y
}
