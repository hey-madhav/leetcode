package main

func uniquePaths(m int, n int) int {
	var dp = make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := range dp {
		dp[i][0] = 1
	}

	for j := range dp[0] {
		dp[0][j] = 1
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
