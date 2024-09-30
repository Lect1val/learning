package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0

	//use two pointer represent for two line in the prob
	for left < right {
		width := right - left

		// min height between left and right is the heighest that water can be contained and not overflow
		minHeight := min(height[left], height[right])

		currentWater := width * minHeight

		if currentWater > maxWater {
			maxWater = currentWater
		}

		//move shortest line inward
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
