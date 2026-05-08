package main

func setZeroes(matrix [][]int) {
	rows := len(matrix)
	columns := len(matrix[0])

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			if matrix[row][column] == 0 { // set flags
				matrix[row][0] = 0 // set first column as zero
				matrix[0][column] = 0 // set first row as zero
			}
		}
	}

	for row := 0; row < rows; row++ {
		if matrix[row][0] == 0 {
			for j:=0;j<columns;j++ {
				matrix[row][j] = 0
			}
		}
	}

	for column := 0;column<columns;columns++ {
		if matrix[0][column] == 0 {
			for i:=0;i<rows;i++ {
				matrix[i][column] = 0
			}
		}
	}
}
