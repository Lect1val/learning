package main

func combinationSum(candidates []int, target int) [][]int {
    var result [][]int
    var combination []int

    var findCombination func(start, target int)
    findCombination = func(start, target int) {
        if target == 0 {
            combinationCopy := make([]int, len(combination))
            copy(combinationCopy, combination)
            result = append(result, combinationCopy)
            return
        }

        for i := start; i < len(candidates); i++ {
            if candidates[i] <= target {
                combination = append(combination, candidates[i])
                findCombination(i, target-candidates[i])
                combination = combination[:len(combination)-1]
            }
        }
    }

    findCombination(0, target)
    
    return result
}
