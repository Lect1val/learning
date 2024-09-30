package main

func canConstruct(ransomNote string, magazine string) bool {
	charCount := make(map[rune]int)

	for _, char := range magazine {
		charCount[char]++
	}

	for _, noteChar := range ransomNote {

		if charCount[noteChar] > 0 {
			charCount[noteChar]--
		} else {
			return false
		}
	}

	return true
}
