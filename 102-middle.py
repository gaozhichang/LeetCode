'''
【中|中】
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

 

示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层序遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

【解题】用临时数组保存当前扫描层
【卡点】
1. python遍历扫描层结果，不能用 for i in tmp，改成 for i in range(len(tmp))
2. 主要扫描到值才添加到结果数组


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
'''

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def levelOrder(self, root):
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """
        # self.dfs(root)
        return self.bfs(root)
    def bfs(self,root):
        #判断为空
        if root is None:
            return []

        #定义变量
        ls = [] #保存当前层节点检索数组
        ls.append(root) #先放入根节点
        res = [] #返回值数组
        while len(ls)>0:
            tmp = ls#取出当前层
            ls = []#清空当前层数数组
            subres = []#当前层的节点val
            #变量当前层的节点
            for i in range(len(tmp)):
                #print(tmp[i].val)
                subres.append(tmp[i].val)#把扫描节点值保存到结果数组
                if tmp[i].left is not None:#如果当前层节点有孩子，保持孩子到检索数组
                    ls.append(tmp[i].left)
                if tmp[i].right is not None:#如果当前层节点有孩子，保持孩子到检索数组
                    ls.append(tmp[i].right)
            #如果读取到了值，保存到结果数组
            if len(subres)>0:
                res.append(subres)
        return res

    def dfs(self,root):
        ls = []
        while root is not None or len(ls)>0:
            if root is not None:
                #中序遍历
                #print(root.val)

                ls.append(root)
                #这里更好riht就是后续遍历
                root=root.left
            else:
                tmp = ls.pop()
                #这输出就是先序遍历
                print(tmp.val)
                root = tmp.right
