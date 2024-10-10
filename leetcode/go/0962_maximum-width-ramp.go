package main

func maxWidthRamp(nums []int) int {
    n := len(nums)

    pairs := make([][2]int,n)
    for i := 0; i < n; i++ {
        pairs[i] = [2]int{nums[i],i}
    }

    sort.Slice(pairs,func(i,j int) bool {
        if pairs[i][0] == pairs[j][0] {
            return pairs[i][1] < pairs[j][1]
        }
        return pairs[i][0] < pairs[j][0]
    })

    maxRamp := 0
    minIndex := n

    for _, pair := range pairs {
        index := pair[1]
        if index > minIndex {
            maxRamp = max(maxRamp, index-minIndex)
        }

        minIndex = min(minIndex, index)
    }

    return maxRamp
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}
