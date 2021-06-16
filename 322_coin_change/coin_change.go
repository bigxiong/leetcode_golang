package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	// return dp1(coins, amount)
	memo = make([]int, amount+1)
	return dp3(coins, amount)
}

func dp1(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	var res int = math.MaxInt32
	for _, coin := range coins {
		subProblem := dp1(coins, amount-coin)
		if subProblem == -1 {
			continue
		}
		res = min(res, 1+subProblem)
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

var memo []int

func dp2(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if memo[amount] != 0 {
		return memo[amount]
	}
	var res int = math.MaxInt32
	for _, coin := range coins {
		subProblem := dp2(coins, amount-coin)
		if subProblem == -1 {
			continue
		}
		res = min(res, 1+subProblem)
	}
	if res == math.MaxInt32 {
		memo[amount] = -1
	} else {
		memo[amount] = res
	}
	return memo[amount]
}

func dp3(coins []int, amount int) int {
	dpTable := make([]int, amount+1)
	for i := 0; i < len(dpTable); i++ {
		dpTable[i] = amount + 1
	}
	dpTable[0] = 0
	for i := 1; i < len(dpTable); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dpTable[i] = min(dpTable[i], 1+dpTable[i-coin])
		}
	}
	if dpTable[amount] == amount+1 {
		return -1
	} else {
		return dpTable[amount]
	}
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

// 打印有多少种凑硬币方法
// 使用coins中硬币（子集中有多少种组合可以凑出amount）
func dp4(coins []int, amount int, start int, result []int) {
	if amount < 0 {
		return
	}
	if amount == 0 {
		fmt.Println(result)
		return
	}
	for i := start; i < len(coins); i++ {
		coin := coins[i]
		result = append(result, coin)
		dp4(coins, amount-coin, i+1, result)
		result = result[:len(result)-1]
	}

}

// 打印有多少种凑硬币方法
// 使用coins中硬币
// 硬币数量不限制，可以重复使用相同面额的。 leetcode 39
func dp5(coins []int, amount int, begin int, result []int) {
	if amount < 0 {
		return
	}
	if amount == 0 {
		fmt.Println(result)
		return
	}
	for i := begin; i < len(coins); i++ {
		coin := coins[i]
		result = append(result, coin)
		dp4(coins, amount-coin, i, result)
		result = result[:len(result)-1]
	}

}

func main() {
	fmt.Println(dp3([]int{1,2,5}, 11))
}
