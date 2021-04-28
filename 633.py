'''
给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c 。

采用数学方法：勾股定理：
首先针对整数的平方和，我们想到的肯定是如何从 a^2 + b^2 = c 这个公式入手，可以考虑的是a从0开始for循环遍历，那么b=Math.sqrt(c-a*a)
题目id：633
12%
12%
'''
import math 
class Solution(object):
    def judgeSquareSum(self, c):
        """
        :type c: int
        :rtype: bool
        """
        res = math.sqrt(c)
        res = math.ceil(res)
        print(res)
        if c==0:
            return True
        for i in range(int(res)):
            find=c-i*i
            if find<0:
                break
            sqrt_find = math.sqrt(find)
            if int(sqrt_find)==sqrt_find:
                print(i,find,int(sqrt_find)==sqrt_find)
                return True
            else:
                i = i+int(sqrt_find)
                continue
                
        return False
