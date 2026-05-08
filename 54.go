package main

func spiralOrder(matrix [][]int) []int {
	left, right := 0, len(matrix)-1
	top, bottom := 0, len(matrix[0])-1

	var result []int

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			result = append(result, matrix[i][top])
		}
		top++

		for i := top; i <= bottom; i++ {
			result = append(result, matrix[right][i])
		}
		right--

		if left <= right {
			for i := right; i >= left; i-- {
				result = append(result, matrix[i][bottom])
			}
			bottom--
		}

		if top<=bottom {
			for i:=bottom;i>=top;i--{
				result = append(result, matrix[left][i])
			}
			left++
		}
	}

	return result
}
