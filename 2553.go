package main

func separateDigits(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	var result []int
	for _, num := range nums {
		if num < 10 {
			result = append(result, num)
		} else {
			result = append(result, convertToSingleDigits(num)...)
		}
	}

	return result
}

func convertToSingleDigits(n int) []int {
	var result []int
	for n/10 > 0 {
		num := n / 10
		n = n % 10
		if num >= 10 {
			result = append(result, convertToSingleDigits(num)...)
		} else {
			result = append(result, num)
		}
	}

	result = append(result, n)

	return result
}
