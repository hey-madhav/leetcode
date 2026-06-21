package main

func createGrid(m int, n int) []string {
	var result []string

	for i := 0; i < m; i++ {
		var s string
		for j := 0; j < n; j++ {
			if i == 0 || j == n-1 {
				s += "."
			} else {
				s += "#"
			}
		}

		result = append(result, s)
	}

	return result
}
