# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(ListNode):
    def __init__(self, val=0, next = None):
        self.current_l1 = None
        self.current_l2 = None
        self.carry = 0

    @property
    def val(self):
        sum = (self.current_l1.val if self.current_l1 else 0) + (self.current_l2.val if self.current_l2 else 0) + self.carry
        if sum < 10:
            self.carry = 0
        else:
            self.carry = 1
            sum -= 10
        return sum

    @property
    def next(self):
        self.current_l1 = self.current_l1.next if self.current_l1 else None
        self.current_l2 = self.current_l2.next if self.current_l2 else None
        if not self.current_l1 and not self.current_l2 and self.carry == 0:
            return None
        return self

    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        self.current_l1 = l1
        self.current_l2 = l2
        return self
