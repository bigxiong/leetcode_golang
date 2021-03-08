package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == -1 {
			return j+1
		}
		if j == -1 {
			return i+1
		}
		if word1[i] == word1[j] {
			return dp(i-1, j-1)
		} else {
			/*
			# 解释：
				# 我直接在 s1[i] 插入一个和 s2[j] 一样的字符
				# 那么 s2[j] 就被匹配了，前移 j，继续跟 i 对比
				# 别忘了操作数加一
			*/
			insert := 1 + dp(i, j-1)


			// word1[i]删掉
			del := 1 + dp(i-1, j)

			/*
			# 解释：
			# 我直接把 s1[i] 替换成 s2[j]，这样它俩就匹配了
			# 同时前移 i，j 继续对比
			# 操作数加一
			*/
			replace := 1 + dp(i-1, j-1)

			return min(min(insert, replace), del)
		}
	}
	return dp(len(word1)-1, len(word2)-1)
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func main() {
	fmt.Println(minDistance("horse","ros"))
}