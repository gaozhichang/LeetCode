/**
给你一棵二叉搜索树，请你 按中序遍历 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。
100%
35.23%
思路：
先取出所有的val（顺序已经是按要求的了），然后再构建树。
卡点：
1.重新构建树，必直接在遍历是构建简单。
2.新构建树时，要用到指针便宜，所以用了go。python不知道怎么用指针
**/


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    vals := []int{}
    var inorder func(*TreeNode)
    inorder = func (root *TreeNode){
        if root!= nil{
            inorder(root.Left)
            vals = append(vals, root.Val)
            inorder(root.Right)
        }
    }
    inorder(root)
    newroot:=&TreeNode{}
    //移动指针，复制右子树
    piontroot := newroot
    for _,i:=range vals{
        piontroot.Right = &TreeNode{Val:i}
        piontroot = piontroot.Right
    }
    return newroot.Right
}
