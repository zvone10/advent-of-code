package aoc2023d20p2

import (
	"fmt"
	"os"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY = 20
const PART = 2

const cutoffTries = 20000

var graph map[string][]string

var flipFlops map[string]bool               // true == is turned on
var conjunctions map[string]map[string]bool // true == high pulse remembered

var inputPeriods map[string][]int

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

	var rxSource string
	for k, v := range graph {
		for _, d := range v {
			if _, ok := conjunctions[d]; ok {
				conjunctions[d][k] = false
			}
		}

		if len(v) == 1 && v[0] == "rx" {
			rxSource = k
		}
	}

	inputPeriods = make(map[string][]int)
	for k, v := range graph {
		if len(v) == 1 && v[0] == rxSource {
			inputPeriods[k] = make([]int, 0)
		}
	}

	queue := make([]signal, 0)

	for i := 0; i < cutoffTries; i++ {
		queue = append(queue, signal{
			module: "broadcaster",
			isLow:  true,
		})
		//lowPulseCount++
		for len(queue) > 0 {
			head := queue[0]
			queue = queue[1:]
			//fmt.Println(head)

			if _, ok := inputPeriods[head.from]; ok && !head.isLow {
				inputPeriods[head.from] = append(inputPeriods[head.from], i)
			}

			if head.module == "broadcaster" {
				dest := graph[head.module]
				for _, d := range dest {
					queue = append(queue, signal{
						module: d,
						from:   "broadcaster",
						isLow:  true,
					})
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
						//highPulseCount++
					}
				} else if head.isLow && v {
					flipFlops[head.module] = false
					for _, d := range dest {
						queue = append(queue, signal{
							module: d,
							from:   head.module,
							isLow:  true,
						})
						//lowPulseCount++
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
						//lowPulseCount++
					} else {
						queue = append(queue, signal{
							module: d,
							from:   head.module,
							isLow:  false,
						})
						//highPulseCount++
					}
				}

			}
		}
	}

	//fmt.Println(lowPulseCount)
	//fmt.Println(highPulseCount)
	fmt.Println(inputPeriods)

	periods := make([]int, 0)

	var total int64
	total = 1

	for _, mp := range inputPeriods {
		p, hasPeriod := checkPeriods(mp)
		if hasPeriod {
			periods = append(periods, p)
			total = lcm(total, int64(p))
		}
	}

	fmt.Println(periods)

	fmt.Println(total)

	internal.PrintResult(DAY, PART, 0)

}

func areAllValuesHigh(m *map[string]bool) bool {
	for _, v := range *m {
		if !v {
			return false
		}
	}

	return true
}

func checkPeriods(numbers []int) (int, bool) {
	p := numbers[1] - numbers[0]
	for i := 2; i < len(numbers); i++ {
		if p != (numbers[i] - numbers[i-1]) {
			return -1, false
		}
	}
	return p, true
}

type signal struct {
	module string
	from   string
	isLow  bool
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int64) int64 {
	return (a * b) / gcd(a, b)
}
