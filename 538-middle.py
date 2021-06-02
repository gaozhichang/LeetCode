'''
[中等|难]
给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。

提醒一下，二叉搜索树满足下列约束条件：

节点的左子树仅包含键 小于 节点键的节点。
节点的右子树仅包含键 大于 节点键的节点。
左右子树也必须是二叉搜索树。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/convert-bst-to-greater-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。」

【解题】
1. 用反先序遍历。
2. 设计一个累加和，边遍历边计算和。并修改节点。
【难点】
1. 左孩子遍历时，没想到怎么把和传递过去。设置一个累加和就行了。
2. python3 可用nonlocal 传递全局变量，python2 得用参数了
2. 树的变量还是不清楚啊。一拳超人学习资料：
https://leetcode-cn.com/problems/convert-bst-to-greater-tree/solution/yi-tao-quan-fa-shua-diao-nge-bian-li-shu-de-wen-5/


'''


# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = 
#         self.left = None
#         self.right = None

class Solution(object):
    def convertBST(self, root):
        """
        :type root: TreeNode
        :rtype: TreeNode
        """
        #遍历得到
        # self.prePost(root)
        total = 0

        def sumRight(total,root): 
            if root is not None:
                total  = sumRight(total,root.right)
                total = total + root.val
                root.val = total
                total  = sumRight(total,root.left)
            return total
        sumRight(total,root)
        
        # 替换
        return root

        

    def prePost(self,root):
        if root is None:
            return
        self.prePost(root.right)
        print(root.val)
        self.prePost(root.left)
