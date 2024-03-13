package reverse_linked_list

func reverse2(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	last := reverse2(head.next)

	head.next.next = head
	head.next = nil

	return last
}
