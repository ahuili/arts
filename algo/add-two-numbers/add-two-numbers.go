package numbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := l1
	n2 := l2
	head := &ListNode{0, nil}
	curr := head
	sum := 0
	carry := 0

	for n1 != nil || n2 != nil {

		sum += carry

		if n1 != nil {
			sum += n1.Val
			n1 = n1.Next
		}

		if n2 != nil {
			sum += n2.Val
			n2 = n2.Next
		}

		carry = sum / 10
		curr.Next = new(ListNode)
		curr.Next.Val = sum % 10
		sum = 0
		curr = curr.Next
	}
	if carry > 0 {
		curr.Next = new(ListNode)
		curr.Next.Val = carry
	}
	return head.Next
}
