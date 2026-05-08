package main

import "math"

func maxPathScore(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])

	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, k+2)

			for c := 0; c < k+2; c++ {
				dp[i][j][c] = math.MinInt
			}
		}
	}

	var dfs func(i, j, c int) int
	dfs = func(i, j, c int) int {
		if c > k || i > m || j > n {
			return math.MinInt
		}

		if i == m-1 && j == n-1 {
			if grid[m-1][n-1] == 0 {
				return 0
			}
			if c < k {
				return grid[m-1][n-1]
			} else {
				return math.MinInt
			}
		}

		if grid[i][j] > 0 {
			c++
		}
		if dp[i][j][c] != math.MinInt {
			return dp[i][j][c]
		}
		r := grid[i][j] + max(dfs(i+1, j, c), dfs(i, j+1, c))
		dp[i][j][c] = r
		return dp[i][j][c]
	}

	r := dfs(0, 0, 0)
	if r < 0 {
		return -1
	}

	return r

}
