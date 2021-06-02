/*
    【简单|稍微难】
    输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

    示例1：

    输入：1->2->4, 1->3->4
    输出：1->1->2->3->4->4

    来源：力扣（LeetCode）
    链接：https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof
    著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

    【解题】
    1. 判断各种情况，两个链表不等长；哪个大等等。
    2. 创建新链表和新链表的当前指针。添加节点移动当前指针。
    【卡点】
    1. 要创建虚拟头节点，要不当前指针会不行。
    2. 当前指针移动赋值。
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    l3:=&ListNode{Val:0,Next:nil}
    next:=l3
    for l1 !=nil || l2!=nil{
        //临时节点
        var tmp ListNode
        if l1!=nil && l2!=nil{
            if l1.Val<=l2.Val{
                tmp = ListNode{Val:l1.Val}
                l1=l1.Next
            }else{
                tmp = ListNode{Val:l2.Val}
                l2=l2.Next
            }
        }else{
            if l1!=nil{
                tmp = ListNode{Val:l1.Val}
                l1=l1.Next
            }
            if l2!=nil{
                tmp = ListNode{Val:l2.Val}
                l2=l2.Next
            }
        }
        //新链表指针赋值和移动
        next.Next = &tmp
        next = next.Next
    }
    //去掉临时头结点
    return l3.Next
}
