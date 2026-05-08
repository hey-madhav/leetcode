package main

func longestPalindrome(s string) string {
	expand := func(left, right int) string {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}

		return s[left+1 : right]
	}

	ans := ""

	for i := 0; i < len(s); i++ {
		odd := expand(i,i)
		if len(odd) > len(ans) {
			ans = odd
		}

		even := expand(i, i+1)
		if len(even) > len(ans) {
			ans = even
		}
	}

	return ans
}
