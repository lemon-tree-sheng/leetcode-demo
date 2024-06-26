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
	if left == nil || left.next == nil {
		return left
	}

	var dummyNode *Node = nil
	p := dummyNode
	q := left

	for q != right.next {
		tmp := q.next
		q.next = p
		p = q
		q = tmp
	}

	return p
}
