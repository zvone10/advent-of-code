package aoc2023d2p1

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
)

const DAY = 2
const PART = 1

var games map[int][]set

func Run() {
	data, err := os.ReadFile("./inputs/2023/d2.txt")
	if err != nil {
		panic(err)
	}

	redExpr, _ := regexp.Compile("[0-9]+ red")
	greenExpr, _ := regexp.Compile("[0-9]+ green")
	//blueExpr, _ := regexp.Compile("[0-9]+ blue")

	numExpr, _ := regexp.Compile("[0-9]+")

	lines := strings.Split(string(data), "\n")
	games = make(map[int][]set)
	for _, l := range lines {
		tokens := strings.Split(l, ": ")
		gt := tokens[0]
		gameNum, _ := strconv.Atoi(strings.Replace(gt, "Game ", "", 1))
		games[gameNum] = make([]set, 0)
		setTokens := strings.Split(tokens[1], "; ")
		for _, st := range setTokens {
			gameSet := set{}
			//fmt.Println(st)
			colors := strings.Split(st, ", ")
			for _, c := range colors {
				if redExpr.Match([]byte(c)) {
					numStr := string(numExpr.Find([]byte(c)))
					num, _ := strconv.Atoi(numStr)
					gameSet.red = num
				} else if greenExpr.Match([]byte(c)) {
					numStr := string(numExpr.Find([]byte(c)))
					num, _ := strconv.Atoi(numStr)
					gameSet.green = num
				} else {
					numStr := string(numExpr.Find([]byte(c)))
					num, _ := strconv.Atoi(numStr)
					gameSet.blue = num
				}
			}

			games[gameNum] = append(games[gameNum], gameSet)
		}
	}
	//fmt.Println(games)
	sum := 0
	for id, game := range games {
		isPossible := true
		for _, s := range game {
			if !(s.red <= 12 && s.green <= 13 && s.blue <= 14) {
				isPossible = false
				break
			}
		}
		if isPossible {
			sum += id
		}
	}

	internal.PrintResult(DAY, PART, sum)
}

type set struct {
	red   int
	blue  int
	green int
}
