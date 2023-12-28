package questions

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	left, right := head, head
	var pre, result *ListNode
	tmpCount := k - 1

	for left != nil {
		for right.Next != nil && tmpCount > 0 {
			right = right.Next
			tmpCount--
		}
		// do reverse
		if tmpCount == 0 {
			next := right.Next
			if pre != nil {
				pre.Next = right
			}
			for i := 0; i < k; i++ {
				tmpNext := left.Next
				left.Next = next
				next = left
				if i == 0 {
					pre = left
				}
				left = tmpNext
			}
			if result == nil {
				result = next
			}

			right = left
			tmpCount = k - 1
			continue
		} else {
			break
		}
	}

	return result
}
