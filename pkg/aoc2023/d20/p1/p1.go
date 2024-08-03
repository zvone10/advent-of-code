package aoc2023d20p1

import (
	"fmt"
	"os"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY = 20
const PART = 1

var graph map[string][]string

var flipFlops map[string]bool               // true == is turned on
var conjunctions map[string]map[string]bool // true == high pulse remembered

func Run() {
	data, err := os.ReadFile("./inputs/2023/d20.txt")
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(data), "\n")

	graph = make(map[string][]string)

	flipFlops = make(map[string]bool)
	conjunctions = make(map[string]map[string]bool)

	for _, line := range inputs {
		tokens := strings.Split(line, " -> ")
		source := tokens[0]
		if source[0] == '&' {
			conjunctions[source[1:]] = make(map[string]bool)
		} else if source[0] == '%' {
			flipFlops[source[1:]] = false
		}

		var sourceName string
		if source == "broadcaster" {
			sourceName = "broadcaster"
		} else {
			sourceName = source[1:]
		}

		graph[sourceName] = strings.Split(tokens[1], ", ")
	}

	for k, v := range graph {
		for _, d := range v {
			if _, ok := conjunctions[d]; ok {
				conjunctions[d][k] = false
			}
		}
	}

	highPulseCount, lowPulseCount := 0, 0
	queue := make([]signal, 0)

	//fmt.Println(graph)

	for i := 0; i < 1000; i++ {
		queue = append(queue, signal{
			module: "broadcaster",
			isLow:  true,
		})
		lowPulseCount++
		for len(queue) > 0 {
			head := queue[0]
			queue = queue[1:]
			//fmt.Println(head)

			if head.module == "broadcaster" {
				dest := graph[head.module]
				for _, d := range dest {
					queue = append(queue, signal{
						module: d,
						from:   "broadcaster",
						isLow:  true,
					})
					lowPulseCount++
				}
			} else if v, ok := flipFlops[head.module]; ok {
				dest := graph[head.module]
				if head.isLow && !v {
					flipFlops[head.module] = !v
					for _, d := range dest {
						queue = append(queue, signal{
							module: d,
							from:   head.module,
							isLow:  false,
						})
						highPulseCount++
					}
				} else if head.isLow && v {
					flipFlops[head.module] = false
					for _, d := range dest {
						queue = append(queue, signal{
							module: d,
							from:   head.module,
							isLow:  true,
						})
						lowPulseCount++
					}
				}
			} else if v, ok := conjunctions[head.module]; ok {
				dest := graph[head.module]
				for _, d := range dest {
					conjunctions[head.module][head.from] = !head.isLow
					if areAllValuesHigh(&v) {
						queue = append(queue, signal{
							module: d,
							from:   head.module,
							isLow:  true,
						})
						//fmt.Println("...")
						lowPulseCount++
					} else {
						queue = append(queue, signal{
							module: d,
							from:   head.module,
							isLow:  false,
						})
						highPulseCount++
					}
				}

			}
		}
	}

	fmt.Println(lowPulseCount)
	fmt.Println(highPulseCount)
	internal.PrintResult(DAY, PART, highPulseCount*lowPulseCount)

}

func areAllValuesHigh(m *map[string]bool) bool {
	for _, v := range *m {
		if !v {
			return false
		}
	}

	return true
}

type signal struct {
	module string
	from   string
	isLow  bool
}
