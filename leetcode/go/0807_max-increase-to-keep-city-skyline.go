package main

func maxIncreaseKeepingSkyline(grid [][]int) int {
    n := len(grid)
    rowMax := make([]int,n)
    colMax := make([]int,n)

    // find and keep the max height of each row and column
    // i represent row
    // j represent column
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] > rowMax[i] {
                rowMax[i] = grid[i][j]
            }
            if grid[i][j] > colMax[j] {
                colMax[j] = grid[i][j]
            }
        }
    }

    increase := 0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            // find minimum height of row and column that each cell can increase the height to reach it
            maxHeight := min(rowMax[i],colMax[j])
            increase +=  maxHeight - grid[i][j]
        }
    }

    return increase
}

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}
