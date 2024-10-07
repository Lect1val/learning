package main

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    charCount := make(map[rune]int)

    for _,char := range s {
        charCount[char]++
    }

    for _, char := range t {
        charCount[char]--

        if  charCount[char] < 0 {
            return false
        }
    }

    return true
}
