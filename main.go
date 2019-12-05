package main

import (
	"fmt"

	"merge-intervals/interval"
)

func main() {
	intervals := []interval.Interval {
		{15, 20},
		{2, 6},
		{17, 24},
	}

	result := interval.Merge(intervals)
	fmt.Println(result)
}
