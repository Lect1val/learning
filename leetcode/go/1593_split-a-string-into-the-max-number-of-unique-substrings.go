package main

func maxUniqueSplit(s string) int {
    return backtrack(s, 0, make(map[string]bool))
}

func backtrack(s string, start int,seen map[string]bool) int {
    if start == len(s) {
        return 0
    }

    maxCount := 0

    for i := start + 1; i <= len(s); i++ {
        substr := s[start:i]
        if !seen[substr] {
            seen[substr] = true
            count := 1 + backtrack(s, i, seen)
            maxCount = max(maxCount, count)

            delete(seen, substr)
        }
    }
    
    return maxCount
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
