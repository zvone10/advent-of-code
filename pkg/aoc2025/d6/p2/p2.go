package aoc2025d6p2

import (
	"fmt"
	"os"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY, PART = 6, 2

var numbers []string
var operators string

func Run() {
	data, err := os.ReadFile("./inputs/2025/d6.txt")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(data), "\n")

	numbers = rows[:4]
	operators = rows[4]
	sum := int64(0)
	currentOperator := ""
	length := 0
	for i, r := range operators {
		if r == ' ' || i == 0 {
			currentOperator += string(r)
			length++
		} else {
			fmt.Println(i, length, currentOperator)
			s0, s1, s2, s3 := numbers[0][i-length:i], numbers[1][i-length:i], numbers[2][i-length:i], numbers[3][i-length:i]
			fmt.Println(s0+"\n", s1+"\n", s2+"\n", s3+"\n", currentOperator+"\n")
			ints := stringsToInts([]string{s0, s1, s2, s3})
			fmt.Println(ints)
			if strings.TrimSpace(currentOperator) == "+" {
				v := int64(0)
				for _, n := range ints {
					v += n
				}
				sum += v
			} else {
				v := int64(1)
				for _, n := range ints {
					v *= n
				}
				sum += v
			}

			length = 1
			currentOperator = string(r)
		}

		if i == len(operators)-1 {
			fmt.Println(i, length, currentOperator)
			s0, s1, s2, s3 := numbers[0][i-length+1:i+1], numbers[1][i-length+1:i+1], numbers[2][i-length+1:i+1], numbers[3][i-length+1:i+1]
			fmt.Println(s0+"\n", s1+"\n", s2+"\n", s3+"\n", currentOperator+"\n")
			ints := stringsToInts([]string{s0, s1, s2, s3})
			fmt.Println(ints)
			if strings.TrimSpace(currentOperator) == "+" {
				v := int64(0)
				for _, n := range ints {
					v += n
				}
				sum += v
			} else {
				v := int64(1)
				for _, n := range ints {
					v *= n
				}
				sum += v
			}
		}

	}

	internal.PrintResult(DAY, PART, sum)
}

func stringsToInts(strs []string) []int64 {
	ints := make([]int64, 0, len(strs))
	for i := len(strs[0]) - 1; i >= 0; i-- {
		n := int64(0)
		for j := 0; j < len(strs); j++ {
			if strs[j][i] == ' ' {
				continue
			} else {
				digit := int64(strs[j][i] - '0')
				n = n*10 + digit
			}
		}
		fmt.Println(n)
		if n == 0 {
			continue
		}
		ints = append(ints, int64(n))
	}
	return ints
}
