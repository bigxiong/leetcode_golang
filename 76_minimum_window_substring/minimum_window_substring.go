package main

import (
	"fmt"
	"math"
	"strings"
)

// 1. brute force
func minWindow1(s string, t string) string {
	var ans string = s
	var ansLen int = math.MaxInt32
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	for i := 0 ; i < len(s); i++ {
		for j := i+len(t); j < len(s)+1; j++ {
			// 检查当前window 是否符合条件
			var valid int
			window := make(map[byte]int)
			//for k := 0 ; k < len(t); k++ {
			//	if strings.ContainsRune(s[i:j], rune(t[k])) {
			//		window[t[k]]++
			//		if window[t[k]] == need[t[k]] {
			//			valid++
			//		}
			//	}
			//}
			for k := i; k < j; k ++ {
					if strings.ContainsRune(t, rune(s[k])) {
						// 当前window中字符在t中出现次数
						window[s[k]]++
						if window[s[k]] == need[s[k]] {
							valid++
						}
					}
			}

			if valid == len(need) && len(s[i:j]) < ansLen {
				ansLen = len(s[i:j])
				ans = s[i:j]
			}
		}
	}
	if ansLen == math.MaxInt32{
		return ""
	}
	return ans
}

func minWindow(s string, t string) string {
	var ans string
	ansLen := math.MaxInt32

	// put t into map
	// 可以用数组代替map
	// var tFreq [128]int
	tFreq := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tFreq[t[i]]++
	}

	windowFreq := make(map[byte]int)
	left := 0
	right := 0
	var valid int
	for right < len(s) {
		c := s[right]
		right++
		// 当前字符出现在T中
		if _, ok := tFreq[c]; ok {
			// 当前窗口中字符c加1
			windowFreq[c]++
			// 出现次数相等，说明改字符已经满足匹配条件
			if windowFreq[c] == tFreq[c] {
				valid++
			}
		}

		// if valid == len(t) {
		// 窗口与T匹配， 缩小窗口
		for valid == len(tFreq) {
			// 更新答案
			if len(s[left:right]) < ansLen {
				ans = s[left:right]
				ansLen = len(s[left:right])
			}
			// 缩小窗口逻辑
			d := s[left]
			// 左侧字符d在T中
			if _, ok := tFreq[d]; ok {
				// 有效匹配字符数减一
				if windowFreq[d] == tFreq[d] {
					valid--
				}
				// 窗口d字符频次减一
				windowFreq[d]--
			}

			// 收缩窗口
			left++
		}
	}

	if ansLen != math.MaxInt32 {
		return ans
	}

	return ""
}

func minWindow3(s string, t string) string {
	var ans string
	ansLen := math.MaxInt32

	// put t into map
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	window := make(map[byte]int)
	left := 0
	right := 0
	var valid int
	for right < len(s) {
		c := s[right]
		right++

		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// if valid == len(t) {
		for valid == len(need) {
			if len(s[left:right]) < ansLen {
				ans = s[left:right]
				ansLen = len(s[left:right])
			}

			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if ansLen != math.MaxInt32 {
		return ans
	}

	return ""
}



func main()  {
	fmt.Println(minWindow1("ADOBECODEBANC", "ABC"))
	// bba  -->  aba
	fmt.Println(minWindow1("bbaa","aba"))
}

