package aoc2025d2p1

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY, PART = 2, 1

func Run() {
	data, err := os.ReadFile("./inputs/2025/d2.txt")
	if err != nil {
		panic(err)
	}

	//digitCounter := make(map[int64]int)
	intervals := strings.Split(string(data), ",")
	sum := int64(0)
	for _, interval := range intervals {
		tokens := strings.Split(interval, "-")
		a, _ := strconv.ParseInt(tokens[0], 10, 64)
		b, _ := strconv.ParseInt(tokens[1], 10, 64)
		fmt.Println(interval)
		for i := a; i <= b; i++ {
			id := strconv.FormatInt(i, 10)
			if id[0:len(id)/2] == id[len(id)/2:] {
				sum += i
			}

		}
	}

	internal.PrintResult(DAY, PART, sum)
}
