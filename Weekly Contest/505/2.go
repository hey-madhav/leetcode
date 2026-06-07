package main

func generateValidStrings(n int, k int) []string {
	var binaries []string

	var dfs func(position int, previousOne bool, cost int, path []byte)
	dfs = func(position int, previousOne bool, cost int, path []byte) {
		if position == n {
			binaries = append(binaries, string(path))
			return
		}

		path = append(path, '0')
		dfs(position+1, false, cost, path)
		path = path[:len(path)-1]

		if !previousOne && cost+position <= k {
			path = append(path, '1')
			dfs(position+1, true, cost+position, path)
			path = path[:len(path)-1]
		}
	}

	dfs(0, false, 0, []byte{})
	return binaries
}
