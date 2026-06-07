package main

func sumOfGoodIntegers(n int, k int) int {
	var totalSum int
	var start = n-k

	if start < 0 {
		start = 1
	}

	for i := start; i <= n+k; i++ {
		if n&i == 0 {
			totalSum += i
		}
	}

	return totalSum
}
