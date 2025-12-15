package aoc2025d6p1

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY, PART = 6, 1

var numbers [][]int
var operators []rune

func Run() {
	data, err := os.ReadFile("./inputs/2025/d6.txt")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(data), "\n")

	re := regexp.MustCompile(` +`)

	for i, row := range rows {
		if i < 4 {
			tokens := re.Split(strings.TrimSpace(row), -1)
			rowNumbers := make([]int, 0, len(tokens))
			for _, t := range tokens {
				n, _ := strconv.Atoi(strings.Trim(t, " "))
				rowNumbers = append(rowNumbers, n)
			}
			numbers = append(numbers, rowNumbers)
		} else {
			operations := re.Split(strings.TrimSpace(row), -1)
			for _, o := range operations {
				trimmed := strings.Trim(o, " ")
				if len(trimmed) == 0 {
					continue
				}
				operators = append(operators, rune(trimmed[0]))
			}
		}
	}

	sum := int64(0)
	for i, o := range operators {
		if o == '+' {
			sum += int64(numbers[0][i] + numbers[1][i] + numbers[2][i] + numbers[3][i])
		} else {
			sum += int64(numbers[0][i] * numbers[1][i] * numbers[2][i] * numbers[3][i])
		}
	}

	internal.PrintResult(DAY, PART, sum)
}
