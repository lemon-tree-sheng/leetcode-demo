package _6_

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	large := &ListNode{}

	smallFlag := small
	largeFlag := large

	flag := head
	for flag != nil {
		if flag.Val < x {
			smallFlag.Next = flag
		} else {
			largeFlag.Next = flag
		}

		flag = flag.Next
	}

	smallFlag.Next = large.Next
	return small.Next
}
