package main

func coinChange(coins []int, amount int) int {
    dp := make([]int,amount+1)
    for i := 1; i <= amount; i++ {
        dp[i] = amount + 1
    }

    dp[0] = 0
    
    for _,coin := range coins {
        for i := coin; i <= amount; i++ {
            if dp[i-coin] +1 < dp[i] {
                dp[i] = dp[i-coin] +1
            }
        }
    }

    if dp[amount] > amount {
        return -1
    }

    return dp[amount]
}
