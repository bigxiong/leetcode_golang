package main

import "sort"

type ArraySlice [][]int

func (p ArraySlice) Len() int           { return len(p) }
// 按end time 排序
func (p ArraySlice) Less(i, j int) bool {
	return p[i][1] < p[j][1]
}

func (p ArraySlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	// 按照结束时间排序
	sort.Sort(ArraySlice(intervals))

	xEnd := intervals[0][1]
	count := 1
	for _, interval := range intervals {
		start := interval[0]
		// 相连，不认为是交叉
		if start >= xEnd {
			count++
			xEnd = interval[1]
		}
	}

	return len(intervals) - count
}

func main() {
	intervals := [][]int{ {1,2}, {2,3}, {3,4}, {1,3}}
	eraseOverlapIntervals(intervals)
}
