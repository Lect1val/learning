package main

func minAddToMakeValid(s string) int {
    left,right := 0,0

    for _,char := range s {
        if char == '(' {
            left++
        } else if char == ')' {
            if left > 0 {
                left--
            } else {
                right++
            }
        }
    }
    
    return left + right
}
