package main


func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	for l := 1; l <= n ; l++ {
		for i := 0; i <= n - l; i++ {
			j := i + l - 1
			if i == j {
				dp[i][j] = 1
				continue
			}

			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	longestPalindromeSubseq("bbbxddjjjbbb")
}
