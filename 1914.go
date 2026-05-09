package main

func rotateGrid(grid [][]int, k int) [][]int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return grid
	}

	m := len(grid)
	n := len(grid[0])
	layers := min(m, n) / 2

	for layer := 0; layer < layers; layer++ {
		top, bottom := layer, m-layer-1
		left, right := layer, n-layer-1

		ring := extractLayer(grid, top, bottom, left, right)
		rotated := rotateRing(ring, k)
		fillLayer(grid, top, bottom, left, right, rotated)
	}

	return grid
}

// extractLayer returns the values around the perimeter of a grid layer in
// traversal order: top row, right column, bottom row, then left column.
func extractLayer(grid [][]int, top, bottom, left, right int) []int {
	var values []int

	for j := left; j <= right; j++ {
		values = append(values, grid[top][j])
	}

	for i := top + 1; i < bottom; i++ {
		values = append(values, grid[i][right])
	}

	if bottom > top {
		for j := right; j >= left; j-- {
			values = append(values, grid[bottom][j])
		}
	}

	for i := bottom - 1; i > top; i-- {
		values = append(values, grid[i][left])
	}

	return values
}

// rotateRing returns a rotated copy of the input slice by k positions.
// The rotation uses modulo arithmetic to avoid unnecessary full-cycle shifts.
func rotateRing(ring []int, k int) []int {
	length := len(ring)
	if length == 0 {
		return nil
	}

	shift := k % length
	if shift == 0 {
		rotated := make([]int, length)
		copy(rotated, ring)
		return rotated
	}

	rotated := make([]int, length)
	for i := 0; i < length; i++ {
		rotated[i] = ring[(i+shift)%length]
	}

	return rotated
}

// fillLayer writes the rotated values back into the perimeter of a grid layer.
func fillLayer(grid [][]int, top, bottom, left, right int, values []int) {
	idx := 0

	for j := left; j <= right; j++ {
		grid[top][j] = values[idx]
		idx++
	}

	for i := top + 1; i < bottom; i++ {
		grid[i][right] = values[idx]
		idx++
	}

	if bottom > top {
		for j := right; j >= left; j-- {
			grid[bottom][j] = values[idx]
			idx++
		}
	}

	for i := bottom - 1; i > top; i-- {
		grid[i][left] = values[idx]
		idx++
	}
}
