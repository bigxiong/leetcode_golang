package main

import (
	"strconv"
	"strings"
)

func multiply(num1, num2 string) string {
	m, n := len(num1), len(num2)
	// 结果最多为 m + n 位数
	var res []int = make([]int, m + n)
	// 从个位数开始逐位相乘
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := (num1[i] - '0') * (num2[j] - '0')
			// 乘积在 res 对应的索引位置
			p1, p2 := i + j, i + j + 1
			// 叠加到 res 上
			sum := int(mul) + res[p2]
			res[p2] = sum % 10
			res[p1] += sum / 10
		}
	}
	// 结果前缀可能存的 0（未使用的位）
	i := 0
	for i < len(res) && res[i] == 0{
		i++
	}
	res = res[i:]
	var valuesText []string
	for i := range res {
		number := res[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	return strings.Join(valuesText, "")
}
