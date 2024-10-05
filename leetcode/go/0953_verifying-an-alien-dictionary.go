package main

func isAlienSorted(words []string, order string) bool {
    //map order of each alien alphabet
    orderMap := make(map[byte]int)
    for i := 0; i < len(order); i++ {
        orderMap[order[i]] = i
    }

    // compare each pair of consecutive words
    for i := 0; i < len(words)-1; i++ {
        if !isCorrectOrder(words[i],words[i+1],orderMap) {
            return false
        }
    }

    return true

}

func isCorrectOrder(word1,word2 string,orderMap map[byte]int) bool {
    //iteration only need iterate up to the lenght of shorter word
    length := len(word1)
    if len(word2) < length {
        length = len(word2)
    }

    //loop through each character in both words up to lenght of shorter word
    //if character from word1 come before word2 return true, if not return false
    for i := 0; i < length; i++ {
        if word1[i] != word2[i] {
            return orderMap[word1[i]] < orderMap[word2[i]]
        }
    }

    //if two words are identical up to the shorter word lenght
    //shorter word should come before the longer word
    return len(word1) <= len(word2)
}

