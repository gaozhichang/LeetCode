'''
【简单|简单】
给定一个整数，写一个函数来判断它是否是 4 的幂次方。如果是，返回 true ；否则，返回 false 。
整数 n 是 4 的幂次方需满足：存在整数 x 使得 n == 4x
【解题】
1. 和求2的幂一样，可以常规位操作（右移动）
2. 因为如果是4的幂数，那么平方根就是2的幂数，所以对平方根求是否是2的幂数，转化成了昨天的题（231）
【卡点】
1.常规很简单。
2.不用循环，想不到开平方（笨啊）
还有其他解法。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/power-of-four
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
'''
class Solution(object):
    def isPowerOfFour(self, n):
        """
        :type n: int
        :rtype: bool
        """


        return self.s2(n)

    def s2(self,n):
        if n<=0:
            return False
        x = int(math.sqrt(n))
        return x*x==n and x&(x-1)==0

    def s1(self,n):
        a = 1
        while a<=n:
            if a==n:
                return True
            a = a<<2
        return False
