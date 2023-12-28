package questions

import "testing"

func Test25(t *testing.T) {
	root := ListNode{Val: 1}
	cur := &root
	cur.Next = &ListNode{Val: 2}
	cur = cur.Next
	cur.Next = &ListNode{Val: 3}
	cur = cur.Next
	cur.Next = &ListNode{Val: 4}
	cur = cur.Next
	cur.Next = &ListNode{Val: 5}
	reverseKGroup(&root, 3)
}
