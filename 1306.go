package main

func canReach(arr []int, start int) bool {
    length := len(arr)

	isSeen := make([]bool, length)

	queue := []int{start}

	for len(queue) > 0 {
		current := queue[0]

		queue = queue[1:]

		if current < 0 || current >= length || isSeen[current] {
			continue
		}

		if arr[current] == 0 {
			return true
		}

		isSeen[current] = true

		queue = append(queue, current + arr[current])
		queue = append(queue, current - arr[current])
	}

	return false
}