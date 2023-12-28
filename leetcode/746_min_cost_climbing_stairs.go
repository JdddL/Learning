package questions

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 1 {
		return cost[0]
	}
	var dp []int = make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i < len(dp); i++ {
		dp[i] = getMin(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(dp)-1]
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
