package main

func minimumTotal(triangle [][]int) int {
	// var dp func(i, j int) int
	// dp = func(i, j int) int {
	//     if i == len(triangle) {
	//         return 0
	//     }
	//     return triangle[i][j] + min(dp(i+1, j), dp(i+1, j+1))
	// }
	// return dp(0, 0)
	m := len(triangle)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	for j := 0; j < len(dp[m-1]); j++ {
		dp[m-1][j] = triangle[m-1][j]
	}
	for i := m - 2; i >= 0; i-- {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = triangle[i][j] + min(dp[i+1][j], dp[i+1][j+1])
		}
	}
	return dp[0][0]
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func main() {
	dp := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	minimumTotal(dp)
}
