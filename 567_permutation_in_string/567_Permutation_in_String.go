package main

func checkInclusion(s1 string, s2 string) bool {
	windowFreq := make(map[byte]int)
	s1Freq := make(map[byte]int)
	for i := 0 ; i < len(s1); i++{
		s1Freq[s1[i]]++
	}
	var left, right int = 0, 0
	var distance int = 0
	for right < len(s2) {
		c := s2[right]
		right++

		// c在s1中出现
		if s1Freq[c] > 0  {
			windowFreq[c]++
			if windowFreq[c] == s1Freq[c] {
				distance++
			}
		}

		for right-left >= len(s1) {
			if distance == len(s1Freq) {
				return true
			}
			c := s2[left]
			left++

			if s1Freq[c] > 0 {
				if windowFreq[c] == s1Freq[c] {
					distance--
				}
				windowFreq[c]--
			}
		}
	}
	return false
}

func main() {
	
}
