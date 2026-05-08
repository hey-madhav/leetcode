package main

import "sort"

func merge(intervals [][]int) [][]int {

    // predefined constant for start (left endpoint), as well as end (right endpoint)
    const START, END = 0, 1
    
    result := make( [][]int, 0);
    
    sort.Slice(intervals, func(a, b int) bool {
        return (intervals[a][0] < intervals[b][0]) || ( (intervals[a][0] == intervals[b][0]) && (intervals[a][1] < intervals[b][1]) )
    })
    
    for _, curInterval := range intervals{
        
        if ( len(result) == 0 ) || ( result[len(result)-1][END] < curInterval[START] ){
            
            // no overlapping
            result = append(result, curInterval)
            
        }else{
            
            // has overlapping
            // merge with previous interval
            result[len(result)-1][END] = max( result[len(result)-1][END], curInterval[END] )
        }
        
    }
    
    return result
}
