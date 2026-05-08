package main

func countSubstrings(s string) int {
    n := len(s)
    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, n)
    }

    count := 0

    // Fill by increasing substring length
    for length := 1; length <= n; length++ {
        for i := 0; i <= n-length; i++ {
            j := i + length - 1

            if s[i] == s[j] && (length <= 2 || dp[i+1][j-1]) {
                dp[i][j] = true
                count++
            }
        }
    }

    return count
}
