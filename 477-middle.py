'''
【中等|难】
两个整数的 汉明距离 指的是这两个数字的二进制数对应位不同的数量。

给你一个整数数组 nums，请你计算并返回 nums 中任意两个数之间汉明距离的总和。
【解题】
1. 按照两两计算搞不定，速度慢
2.官方方法
方法一：逐位统计
在计算汉明距离时，我们考虑的是同一比特位上的值是否不同，而不同比特位之间是互不影响的。

对于数组 \textit{nums}nums 中的某个元素 \textit{val}val，若其二进制的第 ii 位为 11，我们只需统计 \textit{nums}nums 中有多少元素的第 ii 位为 00，即计算出了 \textit{val}val 与其他元素在第 ii 位上的汉明距离之和。

具体地，若长度为 nn 的数组 \textit{nums}nums 的所有元素二进制的第 ii 位共有 cc 个 11，n-cn−c 个 00，则些元素在二进制的第 ii 位上的汉明距离之和为

c\cdot(n-c)
c⋅(n−c)

我们可以从二进制的最低位到最高位，逐位统计汉明距离。将每一位上得到的汉明距离累加即为答案。

具体实现时，对于整数 \textit{val}val 二进制的第 ii 位，我们可以用代码 (val >> i) & 1 来取出其第 ii 位的值。此外，由于 10^9<2^{30}10 
9
 <2 
30
 ，我们可以直接从二进制的第 00 位枚举到第 2929 位。

【难点】
想不到汉明距离结算，可以按位运算，即数第i位有c个1 ，0就是n-c。汉明距离 c(n-c)

'''

class Solution(object):
    def totalHammingDistance(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        total = 0
        n = len(nums)
        for i in range(30):
            #当前第i位的一的数量
            c = 0
            for a in nums:
               c +=  1&a>>i
            #这个位上两两汉明距离总和：c*(n-c) 是1的，和不是1的。
            total += c*(n-c)
        return total
            
    # 这种可以计算任意两个数的汉明距离，但是有点慢要两两计算
    def HammingDistance(self,a,b):
        res = a^b
        # print(res)
        nums = 0
        while res>0:
            nums=nums+(1&res)
            # print(nums,res)
            res=res>>1
        return nums
