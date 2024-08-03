package internal

import (
	"fmt"
	"time"
)

func PrintResult(day uint16, part uint16, result interface{}) {
	fmt.Printf("day: %d, part: %d, result: %v\n", day, part, result)
}

func PrintResultWithTime(day uint16, part uint16, start *time.Time, result interface{}) {
	elapsedTime := time.Since(*start)
	fmt.Printf("day: %d, part: %d, result: %v, execution_time: %d\n", day, part, result, elapsedTime.Milliseconds())
}
