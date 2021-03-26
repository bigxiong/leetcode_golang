package main

// 一次股票买卖，获取最大利润

// 1. brute force
func maxProfit(prices []int) int {
	var max int
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}

// 1. 在最低价格买入，能获得最大利润
func maxProfit2(prices []int) int {
	var minPrice int = prices[0]
	var maxProfit int
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if minPrice != prices[i] {
			if prices[i]-minPrice > maxProfit {
				maxProfit = prices[i] - minPrice
			}
		}
	}
	return maxProfit
}

// 2. 动态规划
func maxProfit3(prices []int) int {
	// L[i] 前i天最低价格
	var minPrices []int = make([]int, len(prices))
	// P[i] 前i天最大利润
	var maxProfits []int = make([]int, len(prices))
	minPrices[0] = prices[0]
	maxProfits[0] = 0
	for i := 1; i < len(prices); i++ {
		minPrices[i] = min(prices[i], minPrices[i-1])
		// 1. 今天操作卖出 2. 今天卖出
		maxProfits[i] = max(prices[i]-minPrices[i], maxProfits[i-1])

	}
	return maxProfits[len(prices)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {

}
