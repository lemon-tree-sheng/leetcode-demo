package reverse_linked_list

func reverseKGroup(head *Node, k int) *Node {
	if head == nil {
		return head
	}

	a := head
	b := head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.next
	}

	newHead := reverseScope(a, b)

	a.next = reverseKGroup(b, k)

	return newHead
}

func reverseScope(left *Node, right *Node) *Node {
	return nil
}
