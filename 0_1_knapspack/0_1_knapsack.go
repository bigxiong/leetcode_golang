package main

// 总容量为W的背包
// N个物品， 第i个物品的重量为wt[i]，价值为val[i]
// 问题: 在满足W的情况下，最能能装多少价值的物品
func knapsack(W, N int, wt, val []int) int{
	dp := make([][]int, N+1)
	for i := 0 ; i < N+1; i++ {
		dp[i] = make([]int, W+1)
	}
	for j := 0; j < W+1; j++ {
		dp[0][j] = 0
	}
	for i := 0; i < N+1; i++ {
		dp[i][0] = 0
	}

	for i := 1; i < N+1; i++ {
		for w := 1; w < W+1; w++ {
			// 重量限制为w, 装入第i个物品或者不装入第i个物品
			if (w - wt[i-1]) >= 0 {
				dp[i][w] = max(dp[i-1][w-wt[i-1]] + val[i-1], dp[i-1][w])
			// 容量不够，第i个物品不能装入
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	return dp[N][W]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// Returns the maximum value that
// can be put in a knapsack of capacity W
// 装第N个物品
func knapSack(W, N int, wt []int, val []int) int {
	// Base Case
	// 当背包容量W为0，或者物品个数N为0时
	if N == 0 || W == 0 {
		return 0
	}


	// If weight of the nth item is more
	// than Knapsack capacity W, then
	// this item cannot be included
	// in the optimal solution
	// 该物品装不下
	if wt[N - 1] > W {
		// 价值等于第N-1个物品价值
		return knapSack(W, N-1, wt, val)

	// Return the maximum of two cases:
	// (1) nth item included
	// (2) not included
	} else {
		return max(val[N - 1]+ knapSack(W - wt[N - 1], N-1, wt, val), knapSack(W, N-1, wt, val))
	}
}

func main() {
	//solution := knapsack(50, 3, []int{10,20,30}, []int{60,100,120})

	solution := knapSack(50, 3, []int{10,20,30}, []int{60,100,120})

	println(solution)
}


