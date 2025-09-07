/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function addTwoNumbers(
    l1: ListNode | null,
    l2: ListNode | null
): ListNode | null {
    const head = new ListNode();
    let current = head;
    let carry = 0;

    while (l1 || l2 || carry) {
        current.next = new ListNode();
        current = current.next;
        let sum = carry + (l1?.val ?? 0) + (l2?.val ?? 0);

        if (sum < 10) {
            carry = 0;
        } else {
            carry = 1;
            sum -= 10;
        }

        current.val = sum;
        l1 = l1?.next;
        l2 = l2?.next;
    }

    return head.next;
}
