package main

var pascalStore [][]int

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}

	pascalStore = append(pascalStore, []int{1})
	pascalStore = append(pascalStore, []int{1, 1})
	for i := 2; i < numRows; i++ {
		pascalArray := []int{}
		pascalArray = append(pascalArray, 1)
		for ind := 1; ind < i; ind++ {
			pascalArray = append(pascalArray, pascalStore[i-1][ind-1]+pascalStore[i-1][ind])
		}
		pascalArray = append(pascalArray, 1)
		pascalStore = append(pascalStore, pascalArray)
	}

	pascalStore = [][]int{}
	return pascalStore
}
