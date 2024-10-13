package main

func topKFrequent(words []string, k int) []string {
    freqMap := make(map[string]int)
    
    for _,word := range words {
        freqMap[word]++
    }

    uniqueWords := make([]string,0,len(freqMap))
    for word := range freqMap {
        uniqueWords = append(uniqueWords, word)
    }

    sort.Slice(uniqueWords, func(i, j int) bool{
        if freqMap[uniqueWords[i]] == freqMap[uniqueWords[j]] {
            return uniqueWords[i] < uniqueWords[j]
        }

        return freqMap[uniqueWords[i]] > freqMap[uniqueWords[j]]
    })

    return uniqueWords[:k]
}