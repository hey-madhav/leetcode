package main

func countSquares(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	rows := len(matrix)
	columns := len(matrix[0])

	var result int
	var squareDimension = 1

	for squareDimension >= min(rows, columns) {
		for i := 0; i < rows;  {
			for j := 0; j < columns; squareDimension++ {

			}
		}
		squareDimension++
	}

	return result
}
