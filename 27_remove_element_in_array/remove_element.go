package main


func removeElement(nums []int, val int) int {

	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			// slow 处是需要删除的元素
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func main() {
	removeElement([]int{3,2,2,3}, 3)
}