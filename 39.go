package main

// combinationSum returns all unique combinations of candidates where the
// candidate numbers sum to target. Each candidate may be used unlimited
// times. The algorithm uses backtracking to explore candidate combinations.
//
// Time complexity:
//   - The algorithm explores combinations recursively, and each path can have
//     up to O(target / min(candidates)) elements in the worst case.
//   - At each level of recursion, it may branch over up to len(candidates)
//     possibilities.
//   - Therefore, the worst-case runtime is exponential in the target value:
//     roughly O(n^(target/minValue)) where n = len(candidates).
//   - This is a loose upper bound for the number of recursive calls.
//
// Space complexity:
//   - The recursion depth is O(target / min(candidates)).
//   - The current combination slice uses O(target / min(candidates)) space.
//   - The result slice stores all valid combinations; if there are k solutions of
//     average length L, that requires O(k * L) additional space.
//   - So total auxiliary space is O(target/minValue + k * L).
func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	backtrack(&result, candidates, []int{}, target, 0)
	return result

}

func backtrack(result *[][]int, candidates, current []int, target, start int) {
	if target < 0 {
		return
	} else if target == 0 {
		// Copy current slice before appending so later backtracking steps do not
		// mutate the stored combination.
		*result = append(*result, append([]int{}, current...))
		return
	}

	for i := start; i < len(candidates); i++ {
		current = append(current, candidates[i])
		backtrack(result, candidates, current, target-candidates[i], i)
		current = current[:len(current)-1]
	}
}
