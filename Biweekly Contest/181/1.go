package main

func validDigit(n int, x int) bool {
	numArray := convertNumToArray(n)
	if len(numArray) == 0 || numArray[len(numArray)-1] == x {
		return false
	}

	for i := 0; i < len(numArray)-1; i++ {
		if numArray[i] == x {
			return true
		}
	}
	return false
}

func convertNumToArray(n int) []int {
	var result []int
	for n > 0 {
		result = append(result, n%10)
		n = n / 10
	}

	return result
}
