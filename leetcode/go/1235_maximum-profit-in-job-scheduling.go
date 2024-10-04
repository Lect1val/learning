package main

func jobScheduling(startTime []int, endTime []int, profit []int) int {
    n:= len(startTime)
    jobs := make([]Job,n)
    
    for i := 0; i < n; i++ {
        jobs[i] = Job{start: startTime[i],end: endTime[i],profit: profit[i]}
    }

    sort.Slice(jobs, func(i,j int) bool {
        return jobs[i].end < jobs[j].end
    })

    dp := make([]int,n+1)
    for i,job := range jobs {
        j := sort.Search(i, func(k int) bool{
            return jobs[k].end > job.start
        })
        dp[i+1] = max(dp[i],dp[j]+job.profit)
    }
    return dp[n]
}

type Job struct {
    start int
    end int
    profit int
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
