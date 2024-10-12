package main

func maximumWealth(accounts [][]int) int {
    maxWealth := 0

    for _,account := range accounts {
        wealth := 0
        for _, bank := range account {
            wealth += bank
        }
        if wealth > maxWealth {
            maxWealth = wealth
        }
    }

    return maxWealth
}
