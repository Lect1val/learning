package main

func numMusicPlaylists(n int, goal int, k int) int {
    const MOD = 1_000_000_007
    dp := make([][]int, goal+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    dp[0][0] = 1 //base case

    for i := 1; i <= goal; i++ {
        for j := 1; j <= n; j++ {
            dp[i][j] = dp[i-1][j-1] * (n - j + 1) % MOD

            if j >k {
                dp[i][j] += dp[i-1][j] * (j - k) % MOD
            }
            dp[i][j] %= MOD
        }
    }

    return dp[goal][n]
}
