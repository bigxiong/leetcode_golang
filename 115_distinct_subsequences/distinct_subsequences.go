package main

func numDistinct(s string, t string) int {
	m := len(t) + 1
	n := len(s) + 1
	// dp[i][j] 代表由t的前i个字符串，可以由s前j字符来产生的方法数
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := i; j < n; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	numDistinct("babgbag", "bag")
}
