/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function (l1, l2) {
    const head = new ListNode(0);
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
};
