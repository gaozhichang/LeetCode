/**
[中|中]
给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。

构造这个链表的 深拷贝。 深拷贝应该正好由 n 个 全新 节点组成，其中每个新节点的值都设为其对应的原节点的值。新节点的 next 指针和 random 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。复制链表中的指针都不应指向原链表中的节点 。

例如，如果原链表中有 X 和 Y 两个节点，其中 X.random --> Y 。那么在复制链表中对应的两个节点 x 和 y ，同样有 x.random --> y 。

返回复制链表的头节点。

用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：

val：一个表示 Node.val 的整数。
random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。
你的代码 只 接受原链表的头节点 head 作为传入参数。

 

示例 1：



输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

【卡点】了解链表的地址链接关系，构造时应该先定义，在挨个赋值。
【解题】
1. 找到原链表随机指针的关系
2. 拷贝新的节点。
3. 然后根据关系挂载next和random

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/copy-list-with-random-pointer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
package main

import "fmt"

type ListNodeRandom struct {
	Val int
	Next *ListNodeRandom
	Random *ListNodeRandom
}

func copyRandomList(head  *ListNodeRandom) *ListNodeRandom {
	rand := make(map[int]int)
	node := make(map[*ListNodeRandom]int)
	newHead := &ListNodeRandom{}
	cur := head
	i:=0
	for cur!=nil{
		//原指针记录地址和索引的关系
		node[cur] = i
		cur = cur.Next
		i++
	}
	cur  = head
	i=0
	for cur!=nil{
		//找到挂载random的索引
		if cur.Random != nil{
			rand[i] = node[cur.Random]
		}
		cur = cur.Next
		i++
	}
	//拼新的链表
	newnode := []*ListNodeRandom{}
	for head!=nil{
		//新的指针节点
		newnode = append(newnode, &ListNodeRandom{Val:head.Val})
		head = head.Next
	}
	cur = newHead
	for i,one:=range newnode{
		if _,ok:=rand[i];ok{
			one.Random = newnode[rand[i]]
		}
		cur.Next = one
		cur = cur.Next
	}
	//fmt.Println(newHead.Next.Val)

	return newHead.Next
}


func main()  {
	fmt.Println(test1Random1())
	//fmt.Println(test2())
}

/**
[[7,null],[13,0],[11,4],[10,2],[1,0]]
 */
func test1Random1() *ListNodeRandom {
	//公共
	var n0,n1,n2,n3,n4 *ListNodeRandom

	n0=&ListNodeRandom{Val:    7,}
	n1=&ListNodeRandom{Val:    13,}
	n2=&ListNodeRandom{Val:    11,}
	n3=&ListNodeRandom{Val:    10,}
	n4=&ListNodeRandom{Val:    1,}

	n0.Next = n1
	n1.Next = n2
	n1.Random = n0
	n2.Next = n3
	n2.Random = n4
	n3.Next = n4
	n3.Random = n2
	n4.Random = n0

	return copyRandomList(n0)
}
