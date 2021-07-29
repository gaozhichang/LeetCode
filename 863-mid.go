

/**
给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 K 。

返回到目标结点 target 距离为 K 的所有结点的值的列表。 答案可以以任何顺序返回。
【中|中]
[卡点]：不知道父亲如何遍历
【解析】：父亲也看成target的一个『孩子』，如果parenthash 向下查找


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/all-nodes-distance-k-in-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    var res []int
    var dfs func (*TreeNode,*TreeNode)
    parent:=make(map[*TreeNode]*TreeNode)
  

    dfs = func(node *TreeNode,pre *TreeNode){
        if node==nil{
            return
        }
        parent[node] = pre
        dfs(node.Left,node)
        dfs(node.Right,node)
    }
    dfs(root,nil)
    
    // fmt.Println(parent)
    //找完父节点，相当于target旁边有三个指针，然后分别递归查找
    var doScan func (*TreeNode,*TreeNode,int)
    doScan=func (node *TreeNode,from *TreeNode,d int){
        if node==nil{
            return
        }
        if d==k{
            res = append(res,node.Val)
            return
        }
        d++
        //递归左字数
        if node.Left!=nil && node.Left!=from{
            doScan(node.Left,node,d)
        }
        //递归右子树。可以不用判断是否为空
        if node.Right!=from{
            doScan(node.Right,node,d)
        }
        //递归父子树
        if parent[node]!=nil && parent[node]!=from{
            doScan(parent[node],node,d)
        }
    }
    
    
    doScan(target,nil,0)

    return res
}
