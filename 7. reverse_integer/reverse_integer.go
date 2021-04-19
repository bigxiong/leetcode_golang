package main

import "math"

func reverse(x int) int {
	var res int
	var tmp = 0
	for x != 0 {
		tmp = x % 10
		if math.MaxInt32/10 < res || (math.MaxInt32/10 == res && math.MaxInt32%10 < tmp) {
			return 0
		}

		if res < math.MinInt32/10 || (math.MinInt32/10 == res && tmp < math.MinInt32%10) {
			return 0
		}

		res = res*10 + tmp
		x = x / 10
	}
	return res
}

func main() {

}
