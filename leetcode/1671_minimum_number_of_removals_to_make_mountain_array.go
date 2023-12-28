package questions

// len(nums)>=3
func minimumMountainRemovals(nums []int) int {
	leftDp, rightDp := make([]int, len(nums)), make([]int, len(nums))
	for i, peak := range nums {
		leftMax := 0
		l := i - 1
		for l >= 0 {
			if peak > nums[l] {
				if leftDp[l] > leftMax {
					leftMax = leftDp[l]
				}
			}
			l -= 1
		}
		leftDp[i] = leftMax + 1
	}
	for i := len(nums) - 1; i >= 0; i-- {
		peak := nums[i]
		rightMax := 0
		r := i + 1
		for r <= len(nums)-1 {
			if peak > nums[r] {
				if rightDp[r] > rightMax {
					rightMax = rightDp[r]
				}
			}
			r += 1
		}
		rightDp[i] = rightMax + 1
	}
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		if leftDp[i] == 1 || rightDp[i] == 1 {
			continue
		}
		tmpLen := leftDp[i] + rightDp[i] - 1
		if maxLen < tmpLen {
			maxLen = tmpLen
		}
	}
	return len(nums) - maxLen
}
