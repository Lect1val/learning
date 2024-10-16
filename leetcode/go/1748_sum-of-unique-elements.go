package main

func sumOfUnique(nums []int) int {
    sum := 0
    numCount := make(map[int]int)
    for _,num := range nums {
        numCount[num]++
    }
    for num,count := range numCount {
        if count == 1 {
            sum += num
        }
    }
    return sum
}
