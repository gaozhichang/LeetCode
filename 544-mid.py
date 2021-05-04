'''
【中等|困难】
你的面前有一堵矩形的、由 n 行砖块组成的砖墙。这些砖块高度相同（也就是一个单位高）但是宽度不同。每一行砖块的宽度之和应该相等。

你现在要画一条 自顶向下 的、穿过 最少 砖块的垂线。如果你画的线只是从砖块的边缘经过，就不算穿过这块砖。你不能沿着墙的两个垂直边缘之一画线，这样显然是没有穿过一块砖的。

给你一个二维数组 wall ，该数组包含这堵墙的相关信息。其中，wall[i] 是一个代表从左至右每块砖的宽度的数组。你需要找出怎样画才能使这条线 穿过的砖块数量最少 ，并且返回 穿过的砖块数量 。

【解题】 49%，54%
【卡点】 没有思路；看的解题。
1.主要是找到思路。即找缝隙出现的次数。
2.hasmap，key代表缝隙未知，value代表出现次数，其实只用到出现其次。
3.特别注意，要去掉墙的左右边界。
'''

class Solution(object):
    def leastBricks(self, wall):
        """
        :type wall: List[List[int]]
        :rtype: int
        """
        dt = {}
        #用map记录缝隙出现的次数
        for i in range(len(wall)):
            presum = 0
            for j in range(len(wall[i])):
                presum = presum + wall[i][j]
                if dt.has_key(presum):
                    dt[presum] = dt[presum] + 1
                else:
                    dt[presum] = 1
        #最后的presum就是但行总宽度，每行都相等。是边界所以去掉
        dt.pop(presum)
        #找到最多的空隙数量，然后用墙的行数减去空隙就是
        ans = 0
        for i  in dt:
            ans = max(dt[i],ans)

        return len(wall)-ans

