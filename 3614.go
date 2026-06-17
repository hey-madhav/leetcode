package main

import "strings"

func processStr2(s string, k int64) byte {
	var result string

	for _, ch := range s {
		switch ch {
		case '*':
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		case '#':
			result += result
		case '%':
			var temp strings.Builder
			for i := len(result) - 1; i >= 0; i-- {
				temp.WriteString(string(result[i]))
			}

			result = temp.String()
		default:
			result += string(ch)
		}
	}

	if k > int64(len(result)) {
		return byte('.')
	}

	return result[k-1]
}
