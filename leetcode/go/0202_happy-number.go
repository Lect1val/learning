package main

func isHappy(n int) bool {
    //tracking if the n already been seen
    seen := make(map[int]bool)

    for n != 1 {
        // if n is already seen, it means we enter the cycle and will never reach 1
        if seen[n] {
            return false
        }

        seen[n] = true
        n = sumOfSqures(n)
    }

    return true
}

func sumOfSqures(num int) int {
    sum := 0
    //calculate each digit
    for num > 0 {
        //extract last digit to be calculate
        digit := num % 10
        sum += digit * digit
        //remove last digit
        num /= 10
    }
    return sum
}
