package main

import (
	"fmt"
)

func generateParenthesis1(n int) []string {
	var res []string
	// left:  左括号个数
	// right: 右括号个数
	var f func(left int, right int, n int, s string)
	f = func(left int, right int, n int, s string)  {
		if left == n && right == n {
			res = append(res, s)
		}
		// 回溯法
		if left < n {
			f(left+1, right, n, s+"(")
		}
		// 只有当左括号个数 > 右括号个数时候，才能放置右括号
		if right < left {
			f(left, right+1, n, s+")")
		}
	}
	f(0,0,n,"")
	return res
}


func generateParenthesis(n int) []string {
	var res []string
	var f func(left int, right int, s string)
	f = func(left int, right int, s string)  {
		if right < left {
			return
		}
		if left < 0 || right < 0 {
			return
		}

		if left == 0 && right == 0 {
			fmt.Println(s)
			res = append(res, s)
			return
		}

		// if left < n {
		// 	f(left+1, right, n, s+"(")
		// }
		// if right < left {
		// 	f(left, right+1, n, s+")")
		// }
		s = s + "("
		f(left-1, right, s)
		s = s[:len(s)-1]

		s = s + ")"
		f(left, right-1, s)
		s = s[:len(s)-1]

	}
	f(n,n,"")
	return res
}

func main() {
	fmt.Println(generateParenthesis(3))
}
