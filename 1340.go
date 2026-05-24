package main

func maxJumps(arr []int, d int) int {
	dp := make([]int, len(arr))

	for i := range dp {
		dp[i] = -1
	}

	var dfs func(int) int
	dfs = func(x int) int {
		if dp[x] != -1 {
			return dp[x]
		}

		ans := 1

		for i := x + 1; i <= min(len(arr)-1, x+d); i++ {
			if arr[i] >= arr[x] {
				break
			}

			ans = max(ans, 1+dfs(i))
		}

		for i := x - 1; i >= max(0, x-d); i-- {
			if arr[i] >= arr[x] {
				break
			}

			ans = max(ans, 1+dfs(i))
		}

		dp[x] = ans

		return ans
	}

	ans := 1

	for i := 0; i < len(arr); i++ {
		ans = max(ans, dfs(i))
	}

	return ans
}
