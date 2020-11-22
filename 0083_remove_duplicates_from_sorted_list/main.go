package _083_remove_duplicates_from_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var res *ListNode

	for res {
		res = *head.Val
		if head.Next != nil {
			if head.Val != head.Next.Val {
				return &ListNode{head.Val, deleteDuplicates(head.Next)}
			} else {
				if head.Next.Next != nil {
					return &ListNode{head.Val, deleteDuplicates(head.Next.Next)}
				}
			}
		}
	}


	return &ListNode{head.Val, nil}
}
