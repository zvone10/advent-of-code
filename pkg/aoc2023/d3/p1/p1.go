package aoc2023d3p1

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY, PART = 3, 1

func Run() {
	data, err := os.ReadFile("./inputs/2023/d3.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	partNumbers := make([]int, 0)

	re := regexp.MustCompile(`\d+`)

	for index, l := range lines {
		matches := re.FindAllStringIndex(l, -1)
		for _, match := range matches {
			n, _ := strconv.Atoi(l[match[0]:match[1]])

			isPart := false
			for i := -1; i <= 1 && !isPart; i++ {
				for j := match[0] - 1; j <= match[1]; j++ {
					if i == 0 && j >= match[0] && j < match[1] {
						continue
					}

					row := index + i
					if row == -1 || row >= len(lines) || j == -1 || j >= len(l) {
						continue
					}

					if lines[row][j] != '.' {
						isPart = true
						break
					}
				}
			}

			if isPart {
				partNumbers = append(partNumbers, n)
			}

		}
	}

	sum := 0
	for _, n := range partNumbers {
		sum += n
	}
	internal.PrintResult(DAY, PART, sum)
}
