package main

func canPlaceFlowers(flowerbed []int, n int) bool {
    eligiblePot := 0

    for i := 0; i < len(flowerbed); i++{
        if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed) - 1 || flowerbed[i+1] == 0) {
            flowerbed[i] = 1
            eligiblePot++
        }

        if eligiblePot >= n {
            return true
        }
    }
    
    return false
}
