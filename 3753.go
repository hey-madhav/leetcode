package main

import "fmt"

func totalWaviness2(num1 int64, num2 int64) int64 {
	type Node struct {
		count int64
		wave  int64
	}

	var solve func(int64) int64

	solve = func(n int64) int64 {
		if n < 0 {
			return 0
		}

		s := []byte(fmt.Sprint(n))

		type State struct {
			position   int
			started    int
			last       int
			secondLast int
		}

		dp := map[State]Node{}

		var dfs func(int, int, int, int, bool) Node

		dfs = func(pos, started, last, secondLast int, tight bool) Node {
			if pos == len(s) {
				return Node{1, 0}
			}

			state := State{pos, started, last, secondLast}

			if !tight {
				if val, ok := dp[state]; ok {
					return val
				}
			}

			limit := 9
			if tight {
				limit = int(s[pos] - '0')
			}

			result := Node{}

			for i := 0; i <= limit; i++ {
				isTight := tight && i == limit

				if started == 0 && i == 0 {
					next := dfs(pos+1, 0, 10, 10, isTight)

					result.count += next.count
					result.wave += next.wave
				} else {
					var add int64 = 0

					if started == 1 && secondLast != 10 {
						if (last > secondLast && last > i) || (last < secondLast && last < i) {
							add = 1
						}
					}

					iSecondLast := 10
					if started == 1 {
						iSecondLast = last
					}

					next := dfs(pos+1, 1, i, iSecondLast, isTight)

					result.count += next.count
					result.wave += next.wave + add*next.count
				}
			}

			if !tight {
				dp[state] = result
			}

			return result
		}

		return dfs(0, 0, 10, 10, true).wave
	}

	return solve(num2) - solve(num1-1)
}
