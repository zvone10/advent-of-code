package aoc2025d3p1

import (
	"os"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY, PART = 3, 1

func Run() {
	data, err := os.ReadFile("./inputs/2025/d3.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	//fmt.Println(lines)

	sum := 0
	for _, l := range lines {
		// process line l
		// map line(string to slice of integers
		numbers := make([]int, 0, len(l))
		for _, ch := range l {
			num := int(ch - '0')
			numbers = append(numbers, num)
		}

		//fmt.Println(numbers)
		maxI1, maxV1 := maxIndexAndValue(numbers)

		if maxI1 == len(numbers)-1 {
			_, maxV2 := maxIndexAndValue(numbers[:maxI1])
			//fmt.Println(maxV2*10 + maxV1)
			sum += (maxV2*10 + maxV1)
		} else {
			_, maxV2 := maxIndexAndValue(numbers[maxI1+1:])
			sum += (maxV1*10 + maxV2)
			//fmt.Println(maxV1*10 + maxV2)
		}

	}

	internal.PrintResult(DAY, PART, sum)
}

func maxIndexAndValue(a []int) (int, int) {
	maxValue, maxIndex := 0, -1
	for i, v := range a {
		if maxValue < v || maxIndex == -1 {
			maxValue = v
			maxIndex = i
		}
	}

	return maxIndex, maxValue
}
