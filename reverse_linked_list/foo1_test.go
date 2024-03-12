package reverse_linked_list

type Node struct {
	val  int
	next *Node
}

func reverse(head *Node) *Node {

	if head == nil {
		return head
	}

	var pre *Node = nil
	cur := head

	for cur != nil {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}

	return pre
}
