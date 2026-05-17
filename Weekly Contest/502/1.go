package main

func isAdjacentDiffAtMostTwo(s string) bool {
	nums := convertToNumArray(s)
	if len(nums) == 0 {
		return false
	}

	for i := 1; i < len(nums); i++ {
		if abs(nums[i]-nums[i-1]) > 2 {
			return false
		}
	}

	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func convertToNumArray(s string) []int {
	if len(s) == 0 {
		return []int{}
	}

	var result []int
	for i := range s {
		c := s[i]
		if c < '0' || c > '9' {
			return []int{}
		}
		result = append(result, int(c-'0'))
	}

	return result
}
