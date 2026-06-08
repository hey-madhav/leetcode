package main

func pivotArray(nums []int, pivot int) []int {
	var smaller, equal, larger []int

	for _, num := range nums {
		if num < pivot {
			smaller = append(smaller, num)
		}

		if num == pivot {
			equal = append(equal, num)
		}

		if num > pivot {
			larger = append(larger, num)
		}
	}

	var result []int
	result = append(result, smaller...)
	result = append(result, equal...)
	result = append(result, larger...)

	return result
}
