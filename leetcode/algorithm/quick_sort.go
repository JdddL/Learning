package algorithm

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) {
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
	quickSort(arr, left, tmpLeft-1)
	quickSort(arr, tmpLeft+1, right)
}
