package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x%10 == 0 && x != 0 {
		return false
	}

	var revertNumber int
	for x > revertNumber {
		revertNumber = revertNumber*10 + x%10
		x = x / 10
	}
	return revertNumber == x || revertNumber/10 == x
}

func main() {

}
