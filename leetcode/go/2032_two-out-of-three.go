package main

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
    count := make(map[int]int)

    addToCount := func(nums []int, seen map[int]bool) {
        for _,num := range nums {
            if !seen[num] {
                count[num]++
                seen[num] = true
            }
        }
    }

    seen1 := make(map[int]bool)
    seen2 := make(map[int]bool)
    seen3 := make(map[int]bool)

    addToCount(nums1, seen1)
    addToCount(nums2, seen2)
    addToCount(nums3, seen3)

    result := []int{}
    for num,value := range count {
        if value >= 2 {
            result = append(result,num)
        }
    }

    return result
}
