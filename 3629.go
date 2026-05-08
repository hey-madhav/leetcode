package main

func minJumps(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	maximum := maxInSlice(nums)
	spf := buildSmallestPrimeFactor(maximum)
	primeToIndices := groupIndicesByPrimeFactors(nums, spf)

	return bfsMinJumps(nums, spf, primeToIndices)
}

// maxInSlice returns the maximum element in the input slice.
func maxInSlice(nums []int) int {
	maximum := nums[0]
	for _, val := range nums {
		if val > maximum {
			maximum = val
		}
	}
	return maximum
}

// buildSmallestPrimeFactor constructs the smallest prime factor (SPF) table for all
// integers in the range [0, maximum]. This table is used to factor values efficiently.
func buildSmallestPrimeFactor(maximum int) []int {
	spf := make([]int, maximum+1)
	for i := range spf {
		spf[i] = i
	}

	for i := 2; i*i <= maximum; i++ {
		if spf[i] == i {
			for j := i * i; j <= maximum; j += i {
				if spf[j] == j {
					spf[j] = i
				}
			}
		}
	}

	return spf
}

// groupIndicesByPrimeFactors builds a mapping from each prime factor to the list of
// indices whose value contains that prime factor.
func groupIndicesByPrimeFactors(nums []int, spf []int) map[int][]int {
	primeToIndices := make(map[int][]int)

	for i, val := range nums {
		for _, prime := range collectDistinctPrimeFactors(val, spf) {
			primeToIndices[prime] = append(primeToIndices[prime], i)
		}
	}

	return primeToIndices
}

// collectDistinctPrimeFactors returns the distinct prime factors of a value.
func collectDistinctPrimeFactors(value int, spf []int) []int {
	var primes []int
	seen := make(map[int]bool)

	for value > 1 {
		prime := spf[value]
		if !seen[prime] {
			primes = append(primes, prime)
			seen[prime] = true
		}
		value /= prime
	}

	return primes
}

// bfsMinJumps performs a BFS from index 0 until reaching the last index. It explores
// adjacent moves and jumps between indices that share a common prime factor.
func bfsMinJumps(nums []int, spf []int, primeToIndices map[int][]int) int {
	n := len(nums)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = -1
	}

	queue := []int{0}
	dist[0] = 0

	for front := 0; front < len(queue); front++ {
		current := queue[front]
		steps := dist[current]

		if current == n-1 {
			return steps
		}

		// Explore neighbor indices.
		for _, next := range []int{current - 1, current + 1} {
			if next >= 0 && next < n && dist[next] == -1 {
				dist[next] = steps + 1
				queue = append(queue, next)
			}
		}

		// Explore all indices that share any prime factor with the current value.
		for _, prime := range collectDistinctPrimeFactors(nums[current], spf) {
			for _, next := range primeToIndices[prime] {
				if dist[next] == -1 {
					dist[next] = steps + 1
					queue = append(queue, next)
				}
			}
			primeToIndices[prime] = nil
		}
	}

	return -1
}
