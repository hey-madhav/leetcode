package main

import "fmt"

func main() {
	nums := []int{1,4,7}
	queries := [][]int{{0,2,1}, {1,1,2}, {0,0,3}}
	fmt.Println(kthRemainingInteger(nums, queries))
}
