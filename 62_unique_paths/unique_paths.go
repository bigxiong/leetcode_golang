package main

import "fmt"

// 自底向上动态递推
func uniquePaths(m int, n int) int {
	var dp [][]int = make([][]int, m)
	for i := 0 ; i < len(dp); i++{
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++{
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	// 动态递推 自底向上
	for i := 1; i < m; i++{
		for j := 1 ; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]

}
// 空间优化1
func uniquePaths3(m int, n int) int {
	// 1. 空间优化 只需保存上一行的状态
	preRow := make([]int, n)
	curRow := make([]int, n)
	for i := 0 ; i < m ; i++ {
		preRow[i] = 1
		curRow[i] = 1
	}

	// 动态递推 自底向上
	for i := 1; i < m; i++{
		for j := 1 ; j < n; j++ {
			curRow[j] = curRow[j-1] + preRow[i]
		}
		copy(preRow, curRow)
	}
	return curRow[n-1]


}
// 空间优化2
func uniquePaths4(m int, n int) int {
	// 2. 空间优化 只需保存上一行的状态
	curRow := make([]int, n)
	for i := 0 ; i < m ; i++ {
		curRow[i] = 1
	}

	// 动态递推 自底向上
	for i := 1; i < m; i++{
		for j := 1 ; j < n; j++ {
			curRow[j] = curRow[j] + curRow[j-1]
		}
	}
	return curRow[n-1]


}

// 自顶向下递归
func uniquePaths2(m int, n int) int {
	memo := make(map[string]int)
	var uniquePathsHelper func(m, n int) int
	uniquePathsHelper = func(m, n int) int {
		if m <=1 || n <= 1 {
			return 1
		}
		key := fmt.Sprintf("%d_%d", m, n)
		if _, isExist := memo[key]; isExist {
			return memo[key]
		}
		right := uniquePathsHelper(m, n-1)
		down := uniquePathsHelper(m-1, n)

		memo[key] = right + down
		return right + down
	}

	return uniquePathsHelper(m, n)
}

func main() {
	
}
