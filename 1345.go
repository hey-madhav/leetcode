package main

func minJumps2(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}

	indexMap := make(map[int][]int)

	for index, value := range arr {
		indexMap[value] = append(indexMap[value], index)
	}

	isVisited := make([]bool, len(arr))

	isVisited[0] = true

	queue := []int{0}

	var steps int

	for len(queue) > 0 {
		queueLength := len(queue)

		for queueLength > 0 {
			current := queue[0]

			queue = queue[1:]

			if current == len(arr) - 1 {
                return steps
            }

			if current-1 >= 0 && !isVisited[current-1] {
				isVisited[current-1] = true
				queue = append(queue, current-1)
			}

			if current+1 < len(arr) && !isVisited[current+1] {
				isVisited[current+1] = true
				queue = append(queue, current+1)
			}

			for _, index := range indexMap[arr[current]] {
				if !isVisited[index] {
					isVisited[index] = true
					queue = append(queue, index)
				}
			}

			delete(indexMap, arr[current])

			queueLength--
		}

		steps++
	}

	return steps
}
