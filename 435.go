package main

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	var start, end = 0, 1

	sort.Slice(intervals, func(i, j int) bool {
		return (intervals[i][start] < intervals[j][start]) || (intervals[i][start] == intervals[j][start] && intervals[i][end] < intervals[j][end])
	})

	count := 1
	currentInterval := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][start] >= currentInterval[end] {
			currentInterval = intervals[i]
			count++
		}
	}

	return len(intervals) - count
}
