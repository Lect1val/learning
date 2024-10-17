package main

func deleteGreatestValue(grid [][]int) int {
    // sort each row in desc
    for i := range grid {
        sort.Slice(grid[i], func(a,b int) bool {
            return grid[i][a] > grid[i][b]
        })
    }

    answer := 0
    n := len(grid[0])

    //iterate through columns to find the max of each column
    for j := 0; j < n; j++ {
        maxVal := 0
        for i := 0; i < len(grid); i++ {
            if grid[i][j] > maxVal {
                maxVal = grid[i][j]
            }
        }
        answer += maxVal
    }

    return answer
}
