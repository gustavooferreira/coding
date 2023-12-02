package solutions

// AddTwoNumbers_V1 adds two linked lists together.
func AddTwoNumbers_V1(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	var currentNode *ListNode

	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		accumulator := 0

		if l1 != nil {
			accumulator += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			accumulator += l2.Val
			l2 = l2.Next
		}

		accumulator += carry

		digit := accumulator % 10
		carry = accumulator / 10

		node := &ListNode{
			Val:  digit,
			Next: nil,
		}

		if currentNode != nil {
			currentNode.Next = node
		}

		currentNode = node

		if result == nil {
			result = currentNode
		}
	}

	return result
}
