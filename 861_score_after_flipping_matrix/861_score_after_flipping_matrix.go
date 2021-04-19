package main

func matrixScore(A [][]int) int {
	if A == nil || len(A) == 0 {
		return 0
	}
	m := len(A)
	n := len(A[0])
	// 将所有行第一位变为1
	for i := 0; i < m; i++ {
		if A[i][0] == 1 {
			continue
		}
		for j := 0; j < n; j++ {
			if A[i][j] == 0 {
				A[i][j] = 1
			} else {
				A[i][j] = 0
			}
		}
	}

	for j := 1; j < n; j++ {
		var zeroCount int
		var oneCount int
		for i := 0; i < m; i++ {
			if A[i][j] == 0 {
				zeroCount++
			} else {
				oneCount++
			}
		}

		if zeroCount > oneCount {
			for i := 0; i < len(A); i++ {
				if A[i][j] == 0 {
					A[i][j] = 1
				} else {
					A[i][j] = 0
				}
			}
		}
	}
	count := 0
	for i := 0; i < len(A); i++ {
		count = count + binaryConversion(A[i])
	}
	return count
}

func binaryConversion(binary []int) int {
	deci := 0
	for i := 0; i < len(binary); i++ {
		deci = deci*2 + binary[i]
	}
	return deci
}

func main() {
	matrixScore([][]int{

		{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0},
	})
}
