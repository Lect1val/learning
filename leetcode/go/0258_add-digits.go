package main

func addDigits(num int) int {
    // Calculate digital root using properties of modulo 9 arithmetic
    return 1 + (num -1) % 9
}
