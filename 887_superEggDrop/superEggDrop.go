package main

import (
	"fmt"
	"math"


)

func superEggDrop(K int, N int) int {
	var dp func(k, n int) int
	var memo [][]int = make([][]int, K+1)
	for i := 0; i < K+1; i++{
		memo[i] = make([]int, N+1)
		for j := 0; j < N+1; j++{
			memo[i][j] = N + 1
		}
	}
	dp = func(k, n int) int {
		if k == 1 {
			return n
		}
		if n == 0 {
			return 0
		}
		if memo[k][n] != N+1 {
			return memo[k][n]
		}

		var res int = math.MaxInt32
		for i := 1; i <= n; i++{
			res = min(res, max(dp(k-1, i-1), dp(k, n-i))+1)
		}
		memo[k][n] = res
		return res
	}

	return dp(K, N)
}

func min(x, y int) int  {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int  {
	if x < y {
		return y
	} else {
		return x
	}
}

func main() {
	fmt.Println(superEggDrop(5, 5000))
}
