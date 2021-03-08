package main

func shipWithinDays(weights []int, D int) int {
	left := getMax(weights)
	right := getSum(weights)
	//for i := left; i <= right; i++ {
	//	if canShipWithinDays(weights, i, D) {
	//		return i
	//	}
	//}
	for left <= right {
		mid := left + (right - left)/2
		if canShipWithinDays(weights, mid, D) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func canShipWithinDays(weights []int, Capacity, D int) bool {
	var i int
	for day := 0; day < D; day++{
		maxCap := Capacity
		for maxCap-weights[i] >=0 {
			maxCap -= weights[i]
			i++
			if i == len(weights) {
				return true
			}
		}
	}
	return false
}

func getMax(weights []int) int {
	if len(weights) == 0 {
		return -1
	}
	var max int = weights[0]
	for i := 1 ; i < len(weights); i++{
		if weights[i] > max {
			max = weights[i]
		}
	}
	return max
}

func getSum(weights []int) int {
	var sum int
	for i := 0 ; i < len(weights) ; i++ {
		sum += weights[i]
	}
	return sum
}