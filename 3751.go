package main

import "strconv"

func totalWaviness(num1 int, num2 int) int {
	ans := 0

	for x := num1; x <= num2; x++ {
		s := strconv.Itoa(x)

		for i := 1; i < len(s)-1; i++ {
			if (s[i] > s[i-1] && s[i] > s[i+1]) ||
				(s[i] < s[i-1] && s[i] < s[i+1]) {
				ans++
			}
		}
	}

	return ans
}
