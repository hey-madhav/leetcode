package main

import "math"

func largestAltitude(gain []int) int {
	var result = math.MinInt
	var total int
	for _, g := range gain {
		result = max(result, total)
		total += g
	}
	
	result = max(result, total)
	return result
}
