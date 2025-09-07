/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode* head = new ListNode();
        ListNode* current = head;
        int carry = 0;

        while (l1 || l2 || carry != 0 ) {
            current->next = new ListNode();
            current = current->next;

            int l1val = 0;
            if (l1) {
                l1val = l1->val;
                l1 = l1->next;
            }

            int l2val = 0;
            if (l2) {
                l2val = l2->val;
                l2 = l2->next;
            }

            int sum = carry + l1val + l2val;
            if (sum < 10) {
                carry = 0;
            } else {
                carry = 1;
                sum -= 10;
            }

            current->val = sum;
        }

        return head->next;
    }
};
