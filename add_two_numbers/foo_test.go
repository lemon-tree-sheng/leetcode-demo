package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1, s2 []*ListNode
	h1 := l1
	for h1 != nil {
		s1 = append(s1, h1)
		h1 = h1.Next
	}
	h2 := l2
	for h2 != nil {
		s2 = append(s2, h2)
		h2 = h2.Next
	}

	flag := 0
	var sum []*ListNode

	if len(s1) > len(s2) {
		j := len(s2) - 1
		for i := len(s1) - 1; i >= 0; i-- {
			tmp := s1[i].Val + s2[j].Val
			j--

			sum = append(sum, &ListNode{
				Val:  tmp%10 + flag,
				Next: nil,
			})

			if tmp >= 10 {
				flag = 1
			} else {
				flag = 0
			}
		}
	}

	return nil
}
