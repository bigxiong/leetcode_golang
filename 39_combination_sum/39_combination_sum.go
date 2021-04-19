package main

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var backtrack func(candidates []int, tracks []int, start, target int)
	backtrack = func(candidates []int, tracks []int, begin, target int) {
		if target < 0 {
			return
		}
		if target == 0 {
			tmp := make([]int, len(tracks))
			copy(tmp, tracks)
			ans = append(ans, tmp)
			return
		}

		for i := begin; i < len(candidates); i++ {
			tracks = append(tracks, candidates[i])
			backtrack(candidates, tracks, i, target-candidates[i])
			tracks = tracks[:len(tracks)-1]
		}
	}

	backtrack(candidates, nil, 0, target)
	return ans
}

func main() {
	combinationSum([]int{2, 3, 6, 7}, 7)
}
