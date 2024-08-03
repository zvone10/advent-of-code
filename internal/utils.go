package internal

import (
	"strconv"
	"strings"
)

func StringToIntSlice(s string, sep string) *[]int {
	tokens := strings.Split(s, sep)
	arr := make([]int, len(tokens))

	for i := 0; i < len(tokens); i++ {
		v, _ := strconv.Atoi(tokens[i])
		arr[i] = v
	}

	return &arr
}
