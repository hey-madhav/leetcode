package main

import "sort"

// maxDistance finds the maximum possible minimum distance between k points on a square's perimeter.
// Points are mapped to a 1D coordinate along the perimeter, then binary search finds the answer.
func maxDistance3(side int, points [][]int, k int) int {
	perimeter := side * 4
	coords := toPerimeterCoords(points, side)
	sort.Slice(coords, func(i, j int) bool { return coords[i] < coords[j] })

	// Binary search on the answer: largest min-distance where k points can fit
	lo, hi, best := int64(1), int64(side), 0
	for lo <= hi {
		mid := (lo + hi) / 2
		if canPlaceKPoints(coords, int64(perimeter), k, mid) {
			best = int(mid)
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return best
}

// toPerimeterCoords maps each 2D point on the square's boundary to a 1D distance
// measured clockwise from the top-left corner (0,0).
//
// The square is traversed in this order:
//
//	Left edge (x=0):    coord = y
//	Bottom edge (y=side): coord = side + x
//	Right edge (x=side):  coord = 3*side - y
//	Top edge (y=0):       coord = 4*side - x
func toPerimeterCoords(points [][]int, side int) []int64 {
	coords := make([]int64, 0, len(points))
	for _, p := range points {
		x, y := p[0], p[1]
		var coord int64
		switch {
		case x == 0:
			coord = int64(y)
		case y == side:
			coord = int64(side + x)
		case x == side:
			coord = int64(side*3 - y)
		default: // y == 0
			coord = int64(side*4 - x)
		}
		coords = append(coords, coord)
	}
	return coords
}

// canPlaceKPoints checks whether k points can be chosen from coords such that
// every adjacent pair (on the perimeter) is at least `minDist` apart.
// It tries each point as the starting point, wrapping around the perimeter.
func canPlaceKPoints(coords []int64, perimeter int64, k int, minDist int64) bool {
	for _, start := range coords {
		// Points within [start, start + perimeter - minDist] are reachable
		// without "lapping" back to the start.
		windowEnd := start + perimeter - minDist
		cur := start

		placed := 1 // we've placed the first point at `start`
		for placed < k {
			// Find the next point at least minDist ahead
			next := lowerBound(coords, cur+minDist)
			if next == len(coords) || coords[next] > windowEnd {
				cur = -1 // can't place enough points from this start
				break
			}
			cur = coords[next]
			placed++
		}

		if cur >= 0 {
			return true // found a valid placement
		}
	}
	return false
}

// lowerBound returns the index of the first element in arr >= target.
// Equivalent to std::lower_bound in C++.
func lowerBound(arr []int64, target int64) int {
	left, right := 0, len(arr)
	for left < right {
		mid := (left + right) / 2
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
