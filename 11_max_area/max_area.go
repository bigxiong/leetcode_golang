package main

func maxArea1(height []int) int {
	var maxArea int
	for i := 0; i < len(height); i++{
		for j := i+1; j < len(height); j++{
			area := min(height[j], height[i]) * (j-i)
			if maxArea < area {
				maxArea = area
			}
		}
	}
	return maxArea
}

func maxArea2(height []int) int {
	var maxArea int
	left, right := 0, len(height)-1
	for left < right {
		area := min(height[left], height[right]) * (right-left)
		if maxArea < area {
			maxArea = area
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func main() {
	
}
