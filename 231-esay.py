'''
【简单|简单】
给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。

如果存在一个整数 x 使得 n == 2x ，则认为 n 是 2 的幂次方。


【解题】
1.直接位运算
2.直接是利用二进制特性，如果n&（n-1)==0
【卡点】
无


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/power-of-two
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
'''

class Solution(object):
    def isPowerOfTwo(self, n):
        """
        :type n: int
        :rtype: bool
        """
        # return self.s1(n)
        return self.s2(n)
            
    # 常规计算
    def s1(self,n):
        res = False
        a = 1
        while a<=n:
            if a==n:
                res = True
            a = a<<1
        return res
    '''
    若 n = 2^xn=2 
    x
    且 xx 为自然数（即 nn 为 22 的幂），则一定满足以下条件：
    恒有 n & (n - 1) == 0，这是因为：
    nn 二进制最高位为 11，其余所有位为 00；
    n - 1n−1 二进制最高位为 00，其余所有位为 11；
    一定满足 n > 0。
    因此，通过 n > 0 且 n & (n - 1) == 0 即可判定是否满足 n = 2^xn=2 
    x
    链接：https://leetcode-cn.com/problems/power-of-two/solution/power-of-two-er-jin-zhi-ji-jian-by-jyd/
    
    '''
    def s2(self,n):
        return n>0 and n&(n-1)==0
