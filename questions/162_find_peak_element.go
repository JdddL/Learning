package questions

func findPeakElement(nums []int) int {
	binarySearch(nums, 0, len(nums)-1)
	return resultIndex
}

var (
	resultIndex int = -1
)

func binarySearch(nums []int, left, right int) {
	if resultIndex != -1 || left > right {
		return
	}
	mid := left + (right-left)/2
	if mid-1 >= 0 && mid+1 <= len(nums)-1 {
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			resultIndex = mid
			return
		}
	}
	binarySearch(nums, left, mid-1)
	binarySearch(nums, mid+1, right)
}
