package aoc2023d1p2

import (
	"os"
	"sort"
	"strings"
	"unicode"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY = 1
const PART = 2

func Run() {
	numbers := []string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	data, err := os.ReadFile("./inputs/2023/d1.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	sum := 0
	for _, l := range lines {
		nums := []number{}
		for i, r := range l {

			if unicode.IsNumber(r) {
				nums = append(nums, number{
					n:     int(r - '0'),
					index: i,
				})
			}
		}

		current := ""
		for ni, nc := range numbers {
			for i, r := range l {
				current = current + string(r)
				if !strings.HasPrefix(nc, current) {
					current = string(r)
				} else if current == nc {
					nums = append(nums, number{
						n:     ni + 1,
						index: i - len(current) + 1,
					})
					current = ""
				}
			}

		}

		sort.Slice(nums, func(i, j int) bool {
			return nums[i].index < nums[j].index
		})

		if len(nums) == 1 {
			sum = sum + 11*nums[0].n
		} else {
			sum = sum + 10*nums[0].n + nums[len(nums)-1].n
		}

	}

	internal.PrintResult(DAY, PART, sum)
}

type number struct {
	n     int
	index int
}
