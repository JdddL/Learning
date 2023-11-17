package questions

import (
	"sort"
)

// Description(https://leetcode.cn/problems/maximum-sum-queries/description/):
// You are given two 0-indexed integer arrays nums1 and nums2, each of length n, and a 1-indexed 2D array queries where queries[i] = [xi, yi].
// For the ith query, find the maximum value of nums1[j] + nums2[j] among all indices j (0 <= j < n), where nums1[j] >= xi and nums2[j] >= yi, or -1 if there is no j satisfying the constraints.
// Return an array answer where answer[i] is the answer to the ith query.

func maximumSumQueries(nums1, nums2 []int, queries [][]int) []int {
	sums := descSums(nums1, nums2)
	ans := make([]int, len(queries))
	for i, query := range queries {
		maxAns := -1
		for _, sum := range *sums {
			if query[0] <= nums1[sum.index] && query[1] <= nums2[sum.index] {
				maxAns = sum.value
				break
			}
		}
		ans[i] = maxAns
	}
	return ans
}

type sum struct {
	index int
	value int
}

type sumsDesc []sum

func (s sumsDesc) Len() int {
	return len(s)
}

func (s sumsDesc) Less(i, j int) bool {
	return s[i].value > s[j].value
}

func (s sumsDesc) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func descSums(nums1, nums2 []int) *[]sum {
	sums := make([]sum, len(nums1))
	for i := 0; i < len(nums1); i++ {
		sums[i].index = i
		sums[i].value = nums1[i] + nums2[i]
	}
	sort.Sort(sumsDesc(sums))
	return &sums
}
