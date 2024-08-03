package aoc2023d19p1

import (
	"os"
	"strconv"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY = 19
const PART = 1

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
var parts [][]int

func Run() {
	data, err := os.ReadFile("./inputs/2023/d19.txt")
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(data), "\n")

	indexMap := map[byte]int{
		'x': x,
		'm': m,
		'a': a,
		's': s,
	}
	workflows = make(map[string][]*rule)
	defaultWorkflows = make(map[string]string)

	rulesParsing := true
	for _, input := range inputs {
		if rulesParsing && len(input) > 0 {
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

		} else if len(input) == 0 {
			rulesParsing = false
		} else if !rulesParsing {
			tokens := strings.Split(input[1:len(input)-1], ",")
			part := make([]int, 4)
			for i, t := range tokens {
				num, _ := strconv.Atoi(t[2:])
				part[i] = num
			}
			parts = append(parts, part)
		}
	}

	sum := 0
	for _, part := range parts {
		currentWorkflowLabel := "in"

		for currentWorkflowLabel != "R" && currentWorkflowLabel != "A" {
			useDefault := true
			currentWorkflowRules := workflows[currentWorkflowLabel]
			for _, r := range currentWorkflowRules {
				if (r.isLarger && part[r.index] > r.value) || (!r.isLarger && part[r.index] < r.value) {
					useDefault = false
					currentWorkflowLabel = r.workflow
					break
				}
			}

			if useDefault {
				currentWorkflowLabel = defaultWorkflows[currentWorkflowLabel]
			}
		}

		if currentWorkflowLabel == "A" {
			sum += (part[0] + part[1] + part[2] + part[3])
		}

	}

	internal.PrintResult(DAY, PART, sum)

}
