package main

import "fmt"

func main() {
	landStartTime := []int{2, 8}
	landDuration := []int{4, 1}
	waterStartTime := []int{6}
	waterDuration := []int{3}
	fmt.Println(earliestFinishTime(landStartTime, landDuration, waterStartTime, waterDuration))
}
