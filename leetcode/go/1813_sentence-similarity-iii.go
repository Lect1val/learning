package main

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
    words1 := strings.Split(sentence1," ")
    words2 := strings.Split(sentence2," ")

    left, right := 0, 0
    n1, n2 := len(words1), len(words2)

    for left < n1 && left < n2 && words1[left] == words2[left] {
        left++
    }

    for right < n1-left && right < n2-left && words1[n1-1-right] == words2[n2-1-right] {
        right++
    }

    return left+right >= n1 || left+right >= n2
}
