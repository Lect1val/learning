package main

func twoSum(nums []int, target int) []int {
	idxMap := make(map[int]int)

	for currentIdx, num := range nums {

		if previousStoredIdx, exists := idxMap[target-num]; exists {
			return []int{previousStoredIdx, currentIdx}
		}

		idxMap[num] = currentIdx

	}

	return nil
}
