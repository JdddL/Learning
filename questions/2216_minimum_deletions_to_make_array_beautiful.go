package questions

// Description(https://leetcode.cn/problems/minimum-deletions-to-make-array-beautiful/description/):
// You are given a 0-indexed integer array nums. The array nums is beautiful if:

// nums.length is even.
// nums[i] != nums[i + 1] for all i % 2 == 0.
// Note that an empty array is considered beautiful.

// You can delete any number of elements from nums. When you delete an element, all the elements to the right of the deleted element will be shifted one unit to the left to fill the gap created and all the elements to the left of the deleted element will remain unchanged.
// Return the minimum number of elements to delete from nums to make it beautiful.

func minDeletion(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	deletion := 0
	i := 0
	for i < len(nums) {
		if (i-deletion)%2 != 0 { // the real index for i is changed as some elements are deleted
			i++
			continue
		}
		j := i + 1
		for j < len(nums) {
			if nums[j] == nums[i] {
				deletion++
				j++
			} else {
				break
			}
		}
		i = j
	}
	if (len(nums)-deletion)%2 != 0 {
		deletion++
	}
	return deletion
}
