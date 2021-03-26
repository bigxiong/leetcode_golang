package main

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	p1, p2 := 1, 2
	for i := 3; i <= n; i++ {
		pi := p1 + p2
		p1 = p2
		p2 = pi
	}
	return p2
}

func main() {

}
