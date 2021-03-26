package main

import "math"

func myAtoi(s string) int {
	var index int = 0
	var sign int32 = 1
	var total int32 = 0

	//1. empty string
	if s == "" {
		return 0
	}
	//2. trim leading space
	for index < len(s) && s[index] == ' ' {
		index++
	}
	if index > len(s) {
		return 0
	}
	//3. handle signs
	if s[index] == '-' {
		sign = -1
		index++
	} else if s[index] == '+' {
		sign = 1
		index++
	}

	//4. convert number
	for index < len(s) {
		digit := s[index] - '0'
		if digit < 0 || digit > 9 {
			break
		}

		if math.MaxInt32/10 < total || (math.MaxInt32/10 == total && math.MaxInt32%10 < digit) {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		total = 10*total + int32(digit)
		index++
	}
	return int(total * sign)
}

func main() {
	myAtoi(" ")
}
