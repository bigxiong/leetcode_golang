package main

func minEatingSpeed(piles []int, H int) int {
	var maxPiles int = piles[0]
	for _, pile := range piles {
		if pile > maxPiles {
			maxPiles = pile
		}
	}
	// brute force
	//for K := 1; K < maxPiles; K++ {
	//	// 以 speed ko'ko H 小时内吃完香蕉
	//	if canFinish(piles, K, H) {
	//		return K
	//	}
	//}

	left := 1
	right := maxPiles - 1
	for left <= right {
		mid := left + (right-left) / 2
		// 以 speed 是否能在 H 小时内吃完香蕉
		if canFinish(piles, mid, H) {
			if mid > 1 && !canFinish(piles,  mid-1, H) {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}


	return left
}

func canFinish(piles []int, K, H int) bool {
	var costHour int
	for i := 0 ; i < len(piles); i++{
		cost := piles[i]/K
		if piles[i] % K != 0 {
			cost += 1
		}
		costHour += cost
	}
	return costHour <= H
}

func main()  {
	minEatingSpeed([]int{312884470}, 968709470)
}