package main

import "math"

func countKthRoots(l int, r int, k int) int {
	var count int

	for i := 0; i <= int(math.Pow(float64(r), 1.0/float64(k))); i++ {
		num := int(math.Pow(float64(i), float64(k)))
		if num >= l && num <= r {
			count++
		}
	}

	return count
}
