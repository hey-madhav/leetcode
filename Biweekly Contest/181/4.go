package main

import (
	"math"
	"sort"
)

type Fenwick struct {
	tree []int
}

func (f *Fenwick) update(idx, val int) {
	for idx < len(f.tree) {
		f.tree[idx] += val
		idx += idx & -idx
	}
}

func (f *Fenwick) query(idx int) int {
	sum := 0
	for idx > 0 {
		sum += f.tree[idx]
		idx -= idx & -idx
	}
	return sum
}

func kthRemainingInteger(nums []int, queries [][]int) []int {
	var evens []int
	for _, v := range nums {
		if v%2 == 0 {
			evens = append(evens, v)
		}
	}
	maxRank := len(evens)
	rankMap := make(map[int]int)
	for i, v := range evens {
		rankMap[v] = i + 1
	}
	f := &Fenwick{tree: make([]int, maxRank+2)}

	blockSize := int(math.Sqrt(float64(len(nums))))
	type query struct {
		l, r, k, idx int
	}
	var qs []query
	for i, q := range queries {
		qs = append(qs, query{l: q[0], r: q[1], k: q[2], idx: i})
	}
	sort.Slice(qs, func(i, j int) bool {
		bi := qs[i].l / blockSize
		bj := qs[j].l / blockSize
		if bi != bj {
			return bi < bj
		}
		return qs[i].r < qs[j].r
	})

	result := make([]int, len(queries))
	currentL, currentR := 0, -1

	add := func(idx int) {
		if nums[idx]%2 == 0 {
			rank := rankMap[nums[idx]]
			f.update(rank, 1)
		}
	}
	remove := func(idx int) {
		if nums[idx]%2 == 0 {
			rank := rankMap[nums[idx]]
			f.update(rank, -1)
		}
	}

	for _, q := range qs {
		for currentL > q.l {
			currentL--
			add(currentL)
		}
		for currentR < q.r {
			currentR++
			add(currentR)
		}
		for currentL < q.l {
			remove(currentL)
			currentL++
		}
		for currentR > q.r {
			remove(currentR)
			currentR--
		}
		low, high := 1, maxRank+q.k+1 // safe upper
		for low < high {
			mid := (low + high) / 2
			val := 2 * mid
			idx := sort.Search(len(evens), func(i int) bool { return evens[i] > val })
			numExcluded := f.query(idx)
			missing := mid - numExcluded
			if missing >= q.k {
				high = mid
			} else {
				low = mid + 1
			}
		}
		result[q.idx] = 2 * low
	}
	return result
}
