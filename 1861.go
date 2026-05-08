package main

func rotateTheBox(boxGrid [][]byte) [][]byte {
	m := len(boxGrid)
	n := len(boxGrid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if boxGrid[i][j] == '#' {
				for k := j + 1; k < n; k++ {
					if boxGrid[i][k] == '#' || boxGrid[i][k] == '*' {
						break
					}

					boxGrid[i][k], boxGrid[i][j] = boxGrid[i][j], boxGrid[i][k]
				}
			}
		}
	}

	var result = make([][]byte, n)
	for i := range result {
		result[i] = make([]byte, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result[j][m-1-i] = boxGrid[i][j]
		}
	}

	return result
}
