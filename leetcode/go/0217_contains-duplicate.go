package main

func containsDuplicate(nums []int) bool {
    if len(nums) <= 1 {
        return false
    }
    
    store := make(map[int]bool)

    for _,num := range nums {
        if store[num] {
            return true
        }
        store[num] = true
    }

    return false
}
