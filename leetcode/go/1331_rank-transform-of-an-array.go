package main

func arrayRankTransform(arr []int) []int {
    sortedArr := append([]int{},arr...)
    sort.Ints(sortedArr)

    rankMap := make(map[int]int)
    rank := 1
    
    for _, num := range sortedArr {
        if _,exists := rankMap[num]; !exists {
            rankMap[num] = rank
            rank++
        }
    }

    for i,num := range arr {
        arr[i] = rankMap[num]
    }

    return arr
}
