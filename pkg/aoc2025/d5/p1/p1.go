package aoc2025d5p1

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY, PART = 5, 1

var intervals []interval

func Run() {
	data, err := os.ReadFile("./inputs/2025/d5.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(data), "\n\n")
	intervalPart, itemsPart := parts[0], parts[1]

	intervalLines := strings.Split(intervalPart, "\n")
	intervals = make([]interval, 0, len(intervalLines))
	for _, line := range intervalLines {
		tokens := strings.Split(line, "-")
		_ = tokens
		a, _ := strconv.ParseInt(tokens[0], 10, 64)
		b, _ := strconv.ParseInt(tokens[1], 10, 64)
		intervals = append(intervals, interval{start: a, end: b})
	}

	counter := 0
	for _, item := range strings.Split(itemsPart, "\n") {
		n, _ := strconv.ParseInt(item, 10, 64)
		fmt.Println(n)
		for _, i := range intervals {
			if n >= i.start && n <= i.end {
				counter++
				break
			}
		}
	}

	internal.PrintResult(DAY, PART, counter)
}

type interval struct {
	start int64
	end   int64
}
