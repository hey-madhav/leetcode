package main

func numberOfSubstrings(s string) int {
	count := 0
	left := 0
	charCount := make([]int, 3)
	n := len(s)

	for right := 0; right < n; right++ {
		charCount[s[right]-'a']++

		for charCount[0] > 0 && charCount[1] > 0 && charCount[2] > 0 {
			count += n - right
			charCount[s[left]-'a']--
			left++
		}
	}

	return count
}
