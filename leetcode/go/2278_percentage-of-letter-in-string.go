package main

func percentageLetter(s string, letter byte) int {
    charCount := 0

    for i,_ := range s {
        if s[i] == letter {
            charCount++
        }
    }

    return (charCount * 100)  / len(s)
}
