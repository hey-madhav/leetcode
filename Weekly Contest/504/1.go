package main

func digitFrequencyScore(n int) int {
	frequencyMap := make(map[int]int)

	for n > 0 {
		key := n % 10
		n = n / 10
		frequencyMap[key]++
	}

	var sum = 0
	for key, value := range frequencyMap {
		sum += key * value
	}

	return sum
}
