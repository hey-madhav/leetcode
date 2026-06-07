package main

func consecutiveSetBits(n int) bool {
	var bits []int

	for n > 0 {
		remainder := n % 2
		bits = append(bits, remainder)
		n /= 2
	}

	lastValue := bits[0]
	var result = false

	for i := 1; i < len(bits); i++ {
		if bits[i] == 1 && bits[i] == lastValue {
			if result {
				return false
			}
			
			result = true
		}

		lastValue = bits[i]
	}

	return result
}
