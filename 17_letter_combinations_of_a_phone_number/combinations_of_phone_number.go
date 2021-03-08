package main

import (
	"fmt"
	"strings"
)

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	phoneKeyboards := map[int][]string {
		2: []string{"a", "b", "c"},
		3: []string{"d", "e", "f"},
		4: []string{"g", "h", "i"},
		5: []string{"j", "k", "l"},
		6: []string{"m", "n", "o"},
		7: []string{"p", "q", "r", "s"},
		8: []string{"t", "u", "v"},
		9: []string{"w", "x", "y", "z"},
	}

	var ans []string
	var digitNums []int
	for _, s := range digits {
		digitNums = append(digitNums, int(s - '0'))
	}

	var backtrack func(digitNums []int, start int, selected []string)
	backtrack = func(digitNums []int, start int, selected []string) {
		if start == len(digitNums) {
			fmt.Println(strings.Join(selected,""))
			ans = append(ans, strings.Join(selected, ""))
			return
		}
		digitNum := digitNums[start]
		strs := phoneKeyboards[digitNum]
		for i := 0 ; i < len(strs); i++{
			// 1. 显示地进行回溯。 此处可以试用string拼接来模拟回溯
			selected = append(selected, strs[i])

			backtrack(digitNums, start+1, selected)

			selected = selected[:len(selected)-1]
		}
	}

	backtrack(digitNums, 0, nil)
	return ans
}

func main() {
	letterCombinations("")
}
