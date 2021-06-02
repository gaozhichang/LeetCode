/*
【简单|简单】
tmp不能直接等于head，这样指针传递会导致head被置空
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var newL *ListNode
    for head!=nil{
        tmp:=ListNode{
            Val:head.Val,
            Next:nil,
        }
        tmp.Next = newL
        newL = &tmp
        head=head.Next
    }

    return newL
}
