package main

import "fmt"

/**
【简单|中】
【卡点】总感觉有简单的方法，但又想不出来。
【解题】
1. 根据题目中的描述「如果一个节点有两个子节点的话，那么该节点的值等于两个子节点中较小的一个」，我们可以知道，对于二叉树中的任意节点 xx，xx 的值不大于其所有子节点的值，因此：
	对于二叉树中的任意节点 xx，xx 的值不大于以 xx 为根的子树中所有节点的值。
	令 xx 为二叉树的根节点，此时我们可以得出结论：二叉树根节点的值即为所有节点中的最小值。
2.因此，我们可以对整棵二叉树进行一次遍历。设根节点的值为 \textit{rootvalue}rootvalue，我们只需要通过遍历，找出严格大于 \textit{rootvalue}rootvalue 的最小值，即为「所有节点中的第二小的值」。



给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为 2 或 0。如果一个节点有两个子节点的话，那么该节点的值等于两个子节点中较小的一个。

更正式地说，root.val = min(root.left.val, root.right.val) 总成立。

给出这样的一个二叉树，你需要输出所有节点中的第二小的值。如果第二小的值不存在的话，输出 -1 。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type TreeNode struct {
   Val int
   Left *TreeNode
   Right *TreeNode
}
func findSecondMinimumValue(root *TreeNode) int {
	res:=-1
	//因为跟紧点就是最小的值，所以找到比根节点大的一个节点就行了
	minval :=root.Val
	var dfs  func(*TreeNode)
	dfs = func(node *TreeNode) {
		//如果node为空，或者res已经找到一个数，并且在找到的比他大了，就退出
		if node==nil || (res!=-1 && node.Val>=res){
			return
		}
		if node.Val>minval{
			res = node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	return res

}


func main()  {
	root:=&TreeNode{
		Val:   2,
		Left:  &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   5,
			Left:  &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(findSecondMinimumValue(root))
}
