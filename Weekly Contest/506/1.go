package main

func checkGoodInteger(n int) bool {
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	sum, squareSum := 0, 0
	for _, digit := range digits {
		sum += digit
		squareSum += digit * digit
	}

	return (squareSum-sum >= 50)
}
