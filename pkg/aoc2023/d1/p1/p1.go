package aoc2023d1p1

import (
	"os"
	"strings"
	"unicode"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY = 1
const PART = 1

func Run() {
	data, err := os.ReadFile("./inputs/2023/d1.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	sum := 0
	for _, l := range lines {
		number := 0
		last := 0
		for _, r := range l {
			if unicode.IsNumber(r) {
				if number == 0 {
					number = int(r - '0')
				} else {
					last = int(r - '0')
				}
			}
		}
		if last == 0 {
			last = number
		}
		number = 10*number + last
		sum += number
	}

	internal.PrintResult(DAY, PART, sum)
}
