package main

func maxProfit(prices []int) int {
	var profit int = 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			profit = prices[i] - prices[i-1]
		}
	}
	return profit
}

func main() {

}
