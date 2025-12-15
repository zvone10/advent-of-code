package aoc2025d2p2

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY, PART = 2, 2

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
			for l := 1; l <= len(id)/2; l++ {
				// check if pattern repeats n times in string using regexp
				if len(id)%l != 0 {
					continue
				}
				pattern := id[:l]
				regexPattern := fmt.Sprintf("^(%s){%d}$", regexp.QuoteMeta(pattern), len(id)/l)
				re := regexp.MustCompile(regexPattern)
				if re.MatchString(id) {
					sum += i
					break
				}
			}

		}
	}

	internal.PrintResult(DAY, PART, sum)
}
