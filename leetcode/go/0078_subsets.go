package main

func subsets(nums []int) [][]int {
    var result [][]int
    backtrack(nums, 0, []int{}, &result)
    return result
}

func backtrack(nums []int, start int, current []int, result *[][]int) {
    *result = append(*result,append([]int{}, current...))

    for i := start; i < len(nums); i++ {
        current = append(current, nums[i])

        backtrack(nums, i+1, current, result)

        current = current[:len(current)-1]
    }
}
