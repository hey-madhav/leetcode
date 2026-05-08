package main

import "sort"

func minOperations(grid [][]int, x int) int {
	var slice []int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			slice = append(slice, grid[i][j])
		}
	}

	if len(slice) == 1 {
		return 1
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	midValue := slice[len(slice)/2]
	var result int

	for i := 0; i < len(slice); i++ {
		difference := abs(slice[i] - midValue)
		if difference % x != 0 {
			return -1
		}
		result += difference/x
	}

	return result
}
