package main

func guessNumber(n int) int {
    left, right := 1, n

    for left <= right {
        mid := left + (right - left) / 2
        result := guess(mid)

        if result == 0 {
            return mid
        } else if result == 1 {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return -1
}
