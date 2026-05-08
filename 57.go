package main

func insert(intervals [][]int, newInterval []int) [][]int {
    var result [][]int
    i := 0
    n := len(intervals)
    
    // Add all intervals that end before newInterval starts
    for i < n && intervals[i][1] < newInterval[0] {
        result = append(result, intervals[i])
        i++
    }
    
    // Merge overlapping intervals
    for i < n && intervals[i][0] <= newInterval[1] {
        newInterval[0] = min(newInterval[0], intervals[i][0])
        newInterval[1] = max(newInterval[1], intervals[i][1])
        i++
    }
    
    // Add the merged interval
    result = append(result, newInterval)
    
    // Add remaining intervals
    for i < n {
        result = append(result, intervals[i])
        i++
    }
    
    return result
}
