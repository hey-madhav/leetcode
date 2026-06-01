package main

import "sort"

func minimumCost(cost []int) int {
	sort.Ints(cost)

	start := 1
	i := len(cost) - 1
	result := 0
	for i >= 0 {
		if start%3 != 0 {
			result += cost[i]
		}

		start++
		i--
	}

	return result
}
