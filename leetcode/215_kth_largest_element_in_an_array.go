package questions

func findKthLargest(nums []int, k int) int {
	kthQuickSort(nums, 0, len(nums)-1, k)
	return nums[len(nums)-k]
}

func kthQuickSort(arr []int, left, right, k int) {
	if left >= right {
		return
	}
	target := arr[left]
	tmpLeft := left
	tmpRight := right
	for tmpLeft < tmpRight {
		for tmpRight > tmpLeft && tmpRight < len(arr) {
			if arr[tmpRight] <= target {
				arr[tmpLeft] = arr[tmpRight]
				break
			}
			tmpRight -= 1
		}
		for tmpLeft < tmpRight && tmpLeft < len(arr) {
			if arr[tmpLeft] > target {
				arr[tmpRight] = arr[tmpLeft]
				break
			}
			tmpLeft += 1
		}
	}
	arr[tmpLeft] = target
	if len(arr)-tmpLeft == k {
		return
	} else if len(arr)-tmpLeft > k {
		kthQuickSort(arr, tmpLeft+1, right, k)
	} else {
		kthQuickSort(arr, left, tmpLeft-1, k)
	}
}
