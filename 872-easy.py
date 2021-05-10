'''
请考虑一棵二叉树上所有的叶子，这些叶子的值按从左到右的顺序排列形成一个 叶值序列 。
【解题】70%，90%
【卡点】
1. 以为不能用树的遍历。
2. 先序，后续，中序遍历应该都可以。
3. 非递归遍历还有有些难度。用两个if ，一个是左孩子入栈，一个是出栈。
'''



# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def leafSimilar(self, root1, root2):
        """
        :type root1: TreeNode
        :type root2: TreeNode
        :rtype: bool
        """
        # 【递归】 取出叶子节点，并且转换成字符串的数组。下面好用join方法。
        # res1 = self.dfs(root1)
        # res2 = self.dfs(root2)

        #【非递归】自己构建堆栈
        res1 = self.depthFirst(root1)
        res2 = self.depthFirst(root2)
        print(res1)
        print(res2)

        #这里直接比较转成的字符串是否相等
        return "-".join(res1) == "-".join(res2)
    

    '''
    递归遍历
    '''
    def dfs(self, root):
        res = []
        if root.left is not None:
           res = res + self.dfs(root.left)

        if root.right is not None:
            res = res + self.dfs(root.right)
        if root.right is None and root.left is None:
            res.append(str(root.val))

        return res

    '''
    非递归
    '''
    # 深度优先遍历
    def depthFirst(self,root):
        ls = []
        res = []
        while len(ls)>0 or root is not None:
            if root is not None:
                # print("l:",root.val)
                ls.append(root)
                root = root.left
            elif len(ls)>0:
                tmp = ls.pop()
                # print(ls)
                if tmp.right is None and tmp.left is None:
                    res.append(str(tmp.val))
                if tmp.right is not None:
                    root=tmp.right
        return res
