package main

func passwordStrength(password string) int {
	seenMap := make(map[byte]bool)
	specialMap := map[byte]bool{
		'!': true,
		'@': true,
		'#': true,
		'$': true,
	}

	score := 0
	for i := range password {
		if !seenMap[password[i]] && (password[i] >= 'a' && password[i] <= 'z') {
			score += 1
			seenMap[password[i]] = true
		}

		if !seenMap[password[i]] && (password[i] >= 'A' && password[i] <= 'Z') {
			score += 2
			seenMap[password[i]] = true
		}

		if !seenMap[password[i]] && (password[i] >= '0' && password[i] <= '9') {
			score += 3
			seenMap[password[i]] = true
		}

		if !seenMap[password[i]] && (specialMap[password[i]]) {
			score += 5
			seenMap[password[i]] = true
		}
	}

	return score
}
