package main

func furthestDistanceFromOrigin(moves string) int {
	var byteMap = make(map[byte]int)

	for i := range moves {
		byteMap[moves[i]]++
	}

	if byteMap['_'] > 0 {
		if byteMap['L'] > byteMap['R'] {
			byteMap['L'] += byteMap['_']
		} else {
			byteMap['R'] += byteMap['_']
		}
	}

	return abs(byteMap['L'] - byteMap['R'])
}
