package main

func minLights(lights []int) int {
	n := len(lights)
	if n == 0 {
		return 0
	}

	diff := make([]int, n+1)
	for i, val := range lights {
		if val > 0 {
			start := max(0, i-val)
			end := min(n-1, i+val)
			diff[start]++
			diff[end+1]--
		}
	}

	dp := make([]bool, n)
	running := 0
	for i := 0; i < n; i++ {
		running += diff[i]
		dp[i] = running > 0
	}

	var result int
	for i := 0; i < n; {
		if dp[i] {
			i++
			continue
		}
		
		j := i
		for j < n && !dp[j] {
			j++
		}
		count := j - i
		result += max(count/2, 1)
		i = j
	}

	return result
}
