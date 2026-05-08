package main

func solveQueries(nums []int, queries []int) []int {
	n := len(nums)

	// Map each value to its list of positions
	positions := make(map[int][]int)
	for i, num := range nums {
		positions[num] = append(positions[num], i)
	}

	// Precompute min circular distance for each index
	answer := make([]int, n)
	for i := range answer {
		answer[i] = -1
	}

	for _, pos := range positions {
		m := len(pos)
		if m < 2 {
			continue
		}

		for i, curr := range pos {
			prev := pos[(i-1+m)%m]
			next := pos[(i+1)%m]

			distPrev := circularDist(curr, prev, n)
			distNext := circularDist(curr, next, n)

			answer[curr] = min(distPrev, distNext)
		}
	}

	// Resolve queries
	result := make([]int, len(queries))
	for i, idx := range queries {
		result[i] = answer[idx]
	}

	return result
}

// circularDist returns the minimum circular distance between two indices
func circularDist(a, b, n int) int {
	d := abs2(a - b)
	return min(d, n-d)
}

func abs2(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
