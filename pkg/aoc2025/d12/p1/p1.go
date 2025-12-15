package aoc2025d12p1

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY, PART = 12, 1

type present []string

func Run() {
	data, err := os.ReadFile("./inputs/2025/d12.txt")
	if err != nil {
		panic(err)
	}

	presentPattern := regexp.MustCompile(`^[\d:]+$`)
	dimensionsPattern := regexp.MustCompile(`^(\d+)x(\d+): `)
	lines := strings.Split(string(data), "\n")
	var presents []present

	counter := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if presentPattern.MatchString(line) {
			present := lines[i+1 : i+4]
			presents = append(presents, present)
		} else if dimensionsPattern.MatchString(line) {
			matches := dimensionsPattern.FindStringSubmatch(line)
			fmt.Println(matches)

			w, _ := strconv.Atoi(matches[1])
			h, _ := strconv.Atoi(matches[2])

			// in line 36x41: 34 13 30 26 30 23 extract rest of numbers
			itemCountsTokens := strings.Split(line[len(matches[0]):], " ")
			//convert all tokens to numbers
			itemCounts := make([]int, len(itemCountsTokens))
			for i, token := range itemCountsTokens {
				itemCounts[i], _ = strconv.Atoi(token)
			}

			// sum of all itemCounts
			sum := 0
			for _, itemCount := range itemCounts {
				sum += itemCount
			}

			if h*w >= sum*9 {
				counter++
			}
		}
	}

	internal.PrintResult(DAY, PART, counter)
}
