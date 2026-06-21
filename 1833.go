package main

import "sort"

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)

	var count = 0
	for _, cost := range costs {
		if cost <= coins {
			count++
			coins -= cost
		}
	}

	return count
}
