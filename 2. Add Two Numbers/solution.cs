/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int val=0, ListNode next=null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */
public class Solution {
    public ListNode AddTwoNumbers(ListNode l1, ListNode l2) {
        var head = new ListNode();
        var current = head;
        int carry = 0;

        while (l1 != null || l2 != null || carry != 0) {
            current.next = new ListNode();
            current = current.next;

            int sum = carry + (l1?.val ?? 0) + (l2?.val ?? 0);
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
}
