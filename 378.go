package main

func kthSmallest(matrix [][]int, k int) int {
    n := len(matrix)
    if n == 0 || k < 1 || k > n*n {
        // Invalid input; adjust based on problem constraints
        return -1
    }
    
    left, right := matrix[0][0], matrix[n-1][n-1]
    
    for left < right {
        mid := left + (right-left)/2
        count := 0
        
        // Count elements <= mid using binary search per row
        for i := 0; i < n; i++ {
            // Find the rightmost column where matrix[i][j] <= mid
            l, r := 0, n-1
            for l <= r {
                m := l + (r-l)/2
                if matrix[i][m] <= mid {
                    l = m + 1
                } else {
                    r = m - 1
                }
            }
            count += l  // l is the number of elements <= mid in this row
        }
        
        if count < k {
            left = mid + 1
        } else {
            right = mid
        }
    }
    
    return left
}