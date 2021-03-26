package main

import "fmt"

// 1. brute force
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var start, maxLen int
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if j-i+1 > maxLen && isValidPalindrome(s, i, j) {
				maxLen = j - i + 1
				start = i
			}
		}
	}
	if maxLen > 0 {
		return s[start : start+maxLen]
	} else {
		return s[:1]
	}

}

func isValidPalindrome(s string, start, end int) bool {
	for i, j := start, end; i < j; {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

// 2. 双指针
func longestPalindrome2(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		s1 := getPalindrome(s, i, i)
		s2 := getPalindrome(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func getPalindrome(s string, i, j int) string {
	for i >= 0 && j < len(s) {
		if s[i] == s[j] {
			i--
			j++
		} else {
			break
		}
	}
	return s[i+1 : j]
}

// 3.动态规划
func longestPalindrome3(s string) string {

	var start, maxLen int
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				dp[i][j] = dp[i+1][j-1]
			}

			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				start = i
			}
		}
	}
	if maxLen > 0 {
		return s[start : start+maxLen]
	} else {
		return s[:1]
	}
}

func main() {
	fmt.Println(longestPalindrome3("acaba"))
}
