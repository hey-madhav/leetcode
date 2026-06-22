package main

func maxNumberOfBalloons(text string) int {
	balloon := []rune{'b', 'a', 'l', 'o', 'n'}
	charMap := make(map[rune]int)

	for _, ch := range text {
		if exists(ch, balloon) {
			charMap[ch]++
		}
	}

	var result int

	for len(charMap) > 0 {
		if charMap['b'] > 0 && charMap['a'] > 0 && charMap['l'] > 1 && charMap['o'] > 1 && charMap['n'] > 0 {
			result++
			charMap['b']--
			charMap['a']--
			charMap['n']--
			charMap['l'] -= 2
			charMap['o'] -= 2
		} else {
			break
		}
	}

	return result
}

func exists(ch rune, arr []rune) bool {
	for _, a := range arr {
		if a == ch {
			return true
		}
	}

	return false
}
