'''
【中等|简单】
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

【解题】两个链表指针同时向后查找。正好符合加法运算逻辑，个、十、百... ，注意要进位
【难点】
1. 两个链表可能不等长，需要构建虚拟链表高位都是0
2. 循环到最后的链表有可能还需要进位
'''

# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        l3 = ListNode(1,None)
        offset1 = l1
        offset2 = l2
        offset3 = l3
        jin = 0
        while offset1 is not None or offset2 is not None:
            # 解决长短不一致问题
            if offset1 is None:
                offset1 = ListNode(0,None)
            if offset2 is None:
                offset2 = ListNode(0,None)

            # 新的链表构建
            offset3.next = ListNode((offset1.val+offset2.val+jin)%10,None)
            offset3 = offset3.next

            # 判断是否需要进一
            if offset1.val+offset2.val+jin>=10:
               jin = 1
            else:
                jin = 0
            offset1 = offset1.next
            offset2 = offset2.next
        
        # 解决最后一位进一问题
        if jin==1:
            offset3.next = ListNode(1,None)

        return l3.next

        
