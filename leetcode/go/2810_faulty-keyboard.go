package main

func finalString(s string) string {
    result := []rune{}

    for _,char := range s {
        if char == 'i' {
            reverse(result)
        } else {
            result = append(result,char)
        }
    }

    return string(result)
}

func reverse(runes []rune) {
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
}
