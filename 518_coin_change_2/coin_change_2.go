package main

// 完全背包问题
func change(amount int, coins []int) int {
	// 行是硬币个数[0,1...N]
	// 列为amount的数[0,1,2...amount]
	var dp [][]int = make([][]int, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
	}
	//
	for i := 0; i < len(dp); i++ {
		dp[i][0] = 1
	}
	for j := 1 ; j < len(dp[0]); j++{
		dp[0][j] = 0
	}

	for i := 1; i <= len(coins); i++ {
		// 表示amount
		for j := 1; j <= amount; j++ {
			var useCoinI int = 0
			if j - coins[i-1] >= 0 {
				useCoinI = dp[i][j-coins[i-1]]
			}

			var notUseCoinI int = dp[i-1][j]
			dp[i][j] = notUseCoinI + useCoinI
		}
	}
	return dp[len(coins)][amount]
}

func main()  {
	change(5, []int{1,2,5})
}