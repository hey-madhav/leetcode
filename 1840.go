package main

import "sort"

func maxBuilding(n int, restrictions [][]int) int {
	restrictions = append(restrictions, []int{1, 0})
	restrictions = append(restrictions, []int{n, n - 1})

	sort.Slice(restrictions, func(i, j int) bool {
		return restrictions[i][0] < restrictions[j][0]
	})

	m := len(restrictions)

	for i := 1; i < m; i++ {
		distance := restrictions[i][0] - restrictions[i-1][0]
		restrictions[i][1] = min(restrictions[i][1], restrictions[i-1][1]+distance)
	}

	for i := m - 2; i >= 0; i-- {
		distance := restrictions[i+1][0] - restrictions[i][0]
		restrictions[i][1] = min(restrictions[i][1], restrictions[i+1][1]+distance)
	}

	var answer int

	for i := 1; i < m; i++ {
		h1 := restrictions[i-1][1]
		h2 := restrictions[i][1]

		distance := restrictions[i][0] - restrictions[i-1][0]

		peak := max(h1, h2) + (distance-abs(h1-h2))/2

		answer = max(peak, answer)
	}

	return answer
}
