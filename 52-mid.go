package main

import "fmt"

/**
【简单|简单】

输入两个链表，找出它们的第一个公共节点。

如下面的两个链表：
在节点 c1 开始相交。

intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3

【解题】先map，在循环第二个
【卡点】难点，map的键值要存*ListNode 即指针地址，不能是val


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type ListNode struct {
	Val int
	Next *ListNode
}



func getIntersectionNode(headA, headB *ListNode) *ListNode {
	find := make(map[*ListNode]bool)
	for headA !=nil{
		find[headA] = true
		headA = headA.Next
	}
	for headB!=nil{
		if _,ok:=find[headB];ok{
			return headB
		}
		headB = headB.Next
	}
	return nil
}
func main()  {
	fmt.Println(test1())
	fmt.Println(test2())
}

func test1() *ListNode {
	//公共
	lc := ListNode{
		Val:  8,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	l1 := ListNode{
		Val:  4,
		Next: &ListNode{
			Val:  1,
			Next: &lc,
		},
	}
	l2 := ListNode{
		Val:  5,
		Next: &ListNode{
			Val:  0,
			Next: &ListNode{
				Val:  1,
				Next: &lc,
			},
		},
	}
	//headA:=ListNode{
	//	Val:  0,
	//	Next: &l1,
	//}
	//headB:=ListNode{
	//	Val:  0,
	//	Next: &l2,
	//}
	return getIntersectionNode(&l1,&l2)
}

func test2() *ListNode {
	//公共
	lc := ListNode{
		Val:  1,
		Next: nil,
	}

	//headA:=ListNode{
	//	Val:  0,
	//	Next: &lc,
	//}
	//headB:=ListNode{
	//	Val:  0,
	//	Next: &lc,
	//}
	return getIntersectionNode(&lc,&lc)
}
