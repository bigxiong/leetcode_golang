package main

import (
	"fmt"
	"strings"
)

/*
Input: ["flower","flow","flight"]
Output: "fl"
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return  ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	prefix := strs[0]
	for i:= 1; i < len(strs); i++ {
		start := 0
		//for ; start < len(prefix) && start < len(strs[i]); start++ {
		//	if strs[i][start] != prefix[start] {
		//		break
		//	}
		//}
		for ; start < len(prefix) && start < len(strs[i]) && strs[i][start] == prefix[start]; {
			start++
		}
		prefix = prefix[:start]
		if prefix == "" {
			return  ""
		}
	}
	return prefix
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return  ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	prefix := strs[0]
	for i:= 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	strs := []string{"flower","flow","flight"}
	fmt.Println(longestCommonPrefix(strs))
}
