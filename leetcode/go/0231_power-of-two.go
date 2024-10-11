package main

func isPowerOfTwo(n int) bool {
    if n <= 0 {
        return false
    }

    // powers of two in binary have exactly one bit set to 1 (e.g., 1 is 0001, 2 is 0010, 4 is 0100, etc.). 
    // subtracting 1 flips all bits after the rightmost 1 (inclusive)
    // so n & (n - 1) will be 0 only for powers of two.
    return n&(n-1) == 0
}
