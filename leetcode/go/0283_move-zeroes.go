package main

func moveZeroes(nums []int)  {
    lastNonZero := 0

    for _,num := range nums {
        if num != 0 {
            nums[lastNonZero] = num
            lastNonZero++
        }
    }

    for i := lastNonZero; i < len(nums); i++ {
        nums[i] = 0
    }
}
