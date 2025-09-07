/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head

	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		current.Next = &ListNode{}
		current = current.Next
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum < 10 {
			carry = 0
		} else {
			carry = 1
			sum -= 10
		}

		current.Val = sum
	}

	return head.Next
}
