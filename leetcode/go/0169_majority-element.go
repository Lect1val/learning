package main

func majorityElement(nums []int) int {
    candidate := nums[0]
    count := 0
    
    for _, num := range nums {
        if count == 0 {
            candidate = num
        }
        if num == candidate {
            count++
        } else {
            count--
        }
    }
    
    return candidate
}