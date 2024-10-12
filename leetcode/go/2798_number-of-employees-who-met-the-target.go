package main

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    empTargetMet :=  0

    for _, hour := range hours {
        if hour >= target {
            empTargetMet++
        }
    }

    return empTargetMet
}
