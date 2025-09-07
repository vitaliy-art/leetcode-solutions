<?php

/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val = 0, $next = null) {
 *         $this->val = $val;
 *         $this->next = $next;
 *     }
 * }
 */
class Solution
{

    /**
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function addTwoNumbers($l1, $l2)
    {
        $head = new ListNode();
        $current = $head;
        $carry = 0;

        while ($l1 || $l2 || $carry) {
            $current->next = new ListNode();
            $current = $current->next;

            $sum = $carry + ($l1?->val ?? 0) + ($l2?->val ?? 0);
            if ($sum < 10) {
                $carry = 0;
            } else {
                $carry = 1;
                $sum -= 10;
            }

            $current->val = $sum;
            $l1 = $l1?->next;
            $l2 = $l2?->next;
        }

        return $head->next;
    }
}
