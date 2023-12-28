package questions

// Description(https://leetcode.cn/problems/maximum-subarray/description/):
// Given an integer array nums, find the subarray with the largest sum, and return its sum.

func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] = nums[i-1] + nums[i]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}
