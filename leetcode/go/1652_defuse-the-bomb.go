package main

func decrypt(code []int, k int) []int {
    n := len(code)
    result := make([]int,n)

    if k == 0 {
        return result
    }

    direction := 1
    if k < 0 {
        direction = -1
        k = -k
    }

    for i := 0; i < n; i++ {
        sum := 0
        for j := 1; j <= k; j++ {
            index := (i +direction*j + n) % n
            sum += code[index]
        }
        result[i] = sum
    }

    return result
}
