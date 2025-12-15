package aoc2025d5p2

import (
	"os"
	"strconv"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY, PART = 5, 2

var intervals []interval

func Run() {
	data, err := os.ReadFile("./inputs/2025/d5.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(data), "\n\n")
	intervalPart := parts[0]

	intervalLines := strings.Split(intervalPart, "\n")
	intervals = make([]interval, 0, len(intervalLines))
	for _, line := range intervalLines {
		tokens := strings.Split(line, "-")
		_ = tokens
		a, _ := strconv.ParseInt(tokens[0], 10, 64)
		b, _ := strconv.ParseInt(tokens[1], 10, 64)
		intervals = append(intervals, interval{start: a, end: b})
	}

	itemCounter := int64(0)
	shouldContinue := true
	for shouldContinue {
		shouldContinue = false
		for i := 0; i < len(intervals); i++ {
			for j := i + 1; j < len(intervals); j++ {
				if intervals[i].Overlaps(intervals[j]) {
					intervals[i] = intervals[i].Union(intervals[j])
					// remove intervals[j]
					intervals = append(intervals[:j], intervals[j+1:]...)
					shouldContinue = true
					j--
				}
			}
		}
	}

	for _, i := range intervals {
		itemCounter += (i.end - i.start + 1)
	}

	internal.PrintResult(DAY, PART, itemCounter)
}

type interval struct {
	start int64
	end   int64
}

func (i interval) Overlaps(other interval) bool {
	iStart := max(i.start, other.start)
	iEnd := min(i.end, other.end)
	return iStart <= iEnd
}

func (i interval) Union(other interval) interval {
	iStart := min(i.start, other.start)
	iEnd := max(i.end, other.end)
	return interval{start: iStart, end: iEnd}
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
