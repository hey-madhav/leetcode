package main

func hasValidPath(grid [][]int) bool {
	rows, columns := len(grid), len(grid[0])

	// The four possible move directions: left, right, up, down.
	directions := [4][2]int{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}

	// For each street type, store which direction(s) it connects to.
	// 0=left, 1=right, 2=up, 3=down.
	streetDirections := map[int][]int{
		1: {0, 1},
		2: {2, 3},
		3: {0, 3},
		4: {1, 3},
		5: {0, 2},
		6: {1, 2},
	}

	// For a valid connection, the next street must accept the opposite direction.
	opposite := [4]int{1, 0, 3, 2}

	// Track which cells we have already visited.
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, columns)
	}

	type Pair struct {
		Row    int
		Column int
	}

	// BFS queue starting from the top-left cell.
	queue := make([]Pair, 0)
	queue = append(queue, Pair{0, 0})
	visited[0][0] = true
	head := 0

	for head < len(queue) {
		current := queue[head]
		head++

		row, column := current.Row, current.Column
		if row == rows-1 && column == columns-1 {
			// Reached bottom-right corner successfully.
			return true
		}

		// Check all directions allowed by the current street.
		for _, direction := range streetDirections[grid[row][column]] {
			nextRow := row + directions[direction][0]
			nextColumn := column + directions[direction][1]

			if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
				continue
			}

			nextType := grid[nextRow][nextColumn]
			ok := false

			// Confirm the next street connects back to this one.
			for _, nextDirection := range streetDirections[nextType] {
				if nextDirection == opposite[direction] {
					ok = true
					break
				}
			}

			if ok {
				visited[nextRow][nextColumn] = true
				queue = append(queue, Pair{nextRow, nextColumn})
			}
		}
	}

	return false
}
