package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		res = ListNode{}
		num = &res
	)

	for l1 != nil || l2 != nil {
		if l1 == nil {
			num.Next = l2
			break
		}

		if l2 == nil {
			num.Next = l1
			break
		}

		if l1.Val < l2.Val {
			num.Next = l1
			l1 = l1.Next
		} else {
			num.Next = l2
			l2 = l2.Next
		}

		num = num.Next
	}

	return res.Next
}
