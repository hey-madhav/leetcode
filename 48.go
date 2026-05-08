package main

func rotate(matrix [][]int) {
	rowItems := make(map[int][]int)
	for i := 0; i < len(matrix); i++ {
		rowArray := []int{}
		for j := 0; j < len(matrix[0]); j++ {
			rowArray = append(rowArray, matrix[i][j])
		}
		rowItems[i] = rowArray
	}

	for j := len(matrix[0]) - 1; j >= 0; j-- {
		ind := len(matrix[0]) - 1 - j
		for i := 0; i < len(matrix); i++ {
			matrix[i][j] = rowItems[ind][i]
		}
	}
}
