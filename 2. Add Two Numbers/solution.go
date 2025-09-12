/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := l1

	for (l1 != nil && l2 != nil && l1 != l2) || (l1 != nil && l1.Val >= 10) {
		if l2 != nil && l1 != l2 {
			l1.Val += l2.Val
		}

		if l1.Val >= 10 {
			l1.Val -= 10
			if l1.Next != nil {
				l1.Next.Val += 1
			} else if l2 != nil && l2.Next != nil {
				l1.Next = l2.Next
				l1.Next.Val += 1
			} else {
				l1.Next = &ListNode{Val: 1}
			}
		}

		if l1.Next != nil {
			l1 = l1.Next
		} else {
			if l2 != nil {
				l1.Next = l2.Next
			}
			break
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	return head
}
