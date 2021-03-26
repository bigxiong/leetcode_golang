package main

func reverseWords(s string) string {
	sArray := []byte(s)

	sb := trimSpaces(sArray)

	// 1.翻转字符串
	reverse(sb, 0, len(sb)-1)

	// 2.翻转每个单词
	reverseEachWord(sb)

	return string(sb)
}

func trimSpaces(sb []byte) []byte {
	left, right := 0, len(sb)-1
	// 去掉字符串开头的空白字符
	for left <= right && sb[left] == ' ' {
		left++
	}

	// 去掉字符串末尾的空白字符
	for left <= right && sb[right] == ' ' {
		right--
	}

	// 将字符串间多余的空白字符去除
	var sbTmp []byte
	for left <= right {
		c := sb[left]
		if c != ' ' {
			sbTmp = append(sbTmp, c)
		} else if sbTmp[len(sbTmp)-1] != ' ' {
			sbTmp = append(sbTmp, c)
		}
		left++
	}
	return sbTmp
}

func reverse(sb []byte, left, right int) {
	for left < right {
		sb[left], sb[right] = sb[right], sb[left]
		left++
		right--
	}
}

func reverseEachWord(sb []byte) {
	start, end := 0, 0
	for start < len(sb) {
		// 循环至单词的末尾
		for end < len(sb) && sb[end] != ' ' {
			end++
		}
		// 翻转单词
		reverse(sb, start, end-1)
		// 更新start，去找下一个单词
		start = end + 1
		end++
	}
}

func main() {
	reverseWords("  hello world  ")
}
