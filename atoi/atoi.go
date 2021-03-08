package main

import (
	"errors"
	"fmt"
	"math"
)
/*
Implement atoi which converts a string to an integer.

The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

Note:

Only the space character ' ' is considered as whitespace character.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.
Example 1:

Input: "42"
Output: 42
Example 2:

Input: "   -42"
Output: -42
Explanation: The first non-whitespace character is '-', which is the minus sign.
             Then take as many numerical digits as possible, which gets 42.
Example 3:

Input: "4193 with words"
Output: 4193
Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.
Example 4:

Input: "words and 987"
Output: 0
Explanation: The first non-whitespace character is 'w', which is not a numerical
             digit or a +/- sign. Therefore no valid conversion could be performed.
Example 5:

Input: "-91283472332"
Output: -2147483648
Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
             Thefore INT_MIN (−231) is returned.
*/
func myAtoi(str string) int {
	for str != "" && str[0] == ' ' {
		str = str[1:]
	}
	if str == "" {
		return 0
	}

	sign := 1
	if str[0] == '-' {
		str = str[1:]
		sign = -1
	} else if str[0] == '+' {
		str = str[1:]
		sign = 1
	} else if str[0] >= '0' && str[0] <= '9' {
		sign = 1
	} else {
		return 0
	}

	for i, s := range str {
		if s < '0' || s > '9' {
			str = str[:i]
			break
		}
	}
	absNum := 0
	for _, b := range str {
		// b - '0' ==> 得到这个字符类型的数字的真实数值的绝对值
		absNum = absNum*10 + int(b-'0')
		// 检查溢出
		switch {
		case sign == 1 && absNum > math.MaxInt32:
			return math.MaxInt32
		// 这里和正数不一样的是，必须和负号相乘，也就是变成负数，否则永远走不到里面
		case sign == -1 && absNum*sign < math.MinInt32:
			return math.MinInt32
		}
	}
	return sign * absNum
}

func atof(str string) (float64, error) {
	for str != "" && str[0] == ' ' {
		str = str[1:]
	}
	if str == "" {
		return 0.0, errors.New("bad input string")
	}
	sign := 1
	if str[0] == '+' {
		str = str[1:]
		sign = 1
	} else if str[0] == '-' {
		str = str[1:]
		sign = -1
	} else if str[0] >= '0' && str[0] <= '9' {
		sign = 1
	}

	var left float64
	for len(str) > 0 && str[0] >= '0' && str[0] <= '9' {
		left = left*10 + float64(str[0] - '0')
		str = str[1:]
	}
	var fraction float64 = 1.0
	var right float64
	if str[0] == '.' {
		str = str[1:]
		for len(str) > 0 && str[0] >= '0' && str[0] <= '9' {
			fraction = fraction/10
			right = right + float64(str[0]-'0')*fraction
			str = str[1:]
		}
	}
	var isPositiveE bool
	var e int
	if len(str) > 0 && str[0] == 'e' {
		str = str[1:]
		if str[0] == '-' {
			isPositiveE = true
			str = str[1:]
		}
		for len(str) > 0 && str[0] >= '0' && str[0] <= '9' {
			e = e*10 + int(str[0] - '0')
			str = str[1:]
		}

	}
	ef := 1.0
	for e > 0 && isPositiveE {
		ef = ef / 10
		e--
	}

	for e > 0 && !isPositiveE {
		ef = ef * 10
		e--
	}

	result := left + right
	if sign == -1 {
		result = -result
	}
	return result*ef, nil
}


func main() {
	fmt.Println(myAtoi(" "))
	fmt.Println(atof("19.8e10"))
}
