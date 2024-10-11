package main

func dailyTemperatures(temperatures []int) []int {
    result := make([]int, len(temperatures))
    stack := []int{}

    for i, temp := range temperatures {
        // process stack untill finding a cooler temp so that mean this i temp is warmmer
        for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
            prevIndex := stack[len(stack)-1]
            stack = stack[:len(stack)-1] // pop the top of stack
            result[prevIndex] = i - prevIndex // store the result for this prev index
        }

        stack = append(stack,i)
    }

    return result
}
