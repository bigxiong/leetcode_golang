package main

import "math"

func lengthOfLongestSubstring(s string) int {
	windowFreq := make(map[byte]int)
	left, right := 0, 0
	var longestLen int = math.MinInt32
	for right < len(s) {
		c := s[right]
		right++

		windowFreq[c]++
		// 收缩窗口
		//if windowFreq[c] > 1 {
		for windowFreq[c] > 1 {
			windowFreq[s[left]]--
			left++
		}
		if (right - left) > longestLen {
			longestLen = right - left
		}
	}
	return longestLen
}

func main() {
	lengthOfLongestSubstring("pwwkew")
}