package main

func maxDistance(colors []int) int {
	indexMap := make(map[int][]int)
	colorArray := make([]int, 0)
	for i, color := range colors {
		if _, ok := indexMap[color]; !ok {
			colorArray = append(colorArray, color)
		}
		indexMap[color] = append(indexMap[color], i)
	}

	var maxDistance int
	for i := 0; i < len(colorArray)-1; i++ {
		color1Distance := indexMap[colorArray[i]][0]
		color1Distance2 := indexMap[colorArray[i]][len(indexMap[colorArray[i]])-1]
		for j := i + 1; j < len(colorArray); j++ {
			color2Distance := indexMap[colorArray[j]][len(indexMap[colorArray[j]])-1]
			color2Distance2 := indexMap[colorArray[j]][0]
			maxDistance = max(maxDistance, abs(color2Distance - color1Distance))
			maxDistance = max(maxDistance, abs(color2Distance2 - color1Distance2))
		}
	}

	return maxDistance
}
