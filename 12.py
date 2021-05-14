'''
【中等|中等】
罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。 
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给你一个整数，将其转为罗马数字。
【解题】
1. 逐位进行转换，高位到低位
2. 单个数字转换，区分多种情况。
【卡点】
单数字转换多种情况判断。要考虑全面
'''

class Solution(object):
    def intToRoman(self, num):
        """
        :type num: int
        :rtype: str
        """
        self.luomaHash = {0:"",1:"I",5:"V",10:"X",50:"L",100:"C",500:"D",1000:"M"}
        
        res = []
        w  = 1000
        # 拆分数组，逐个解决
        while w>=1:
#             print("--xx->",num%w)
            if num%w==num:
                w = int(w/10)
#                 print("--->",w,num)
            else:
                k = int(num/w)
#                 print("---1>",k,w,num)
                res.append(self.trans(k,w))
                num = num%w
        res.append(self.trans(num,1))
        # print(res)
        return "".join(res)
        
    def trans(self, k,w):
#         print(k,w)
        # 分层四中情况：
        # 1. 余数等于4，又分为：除以5>0, 除以5<1
        # 2. 余数不等于4
        if k%5==4:
            if int(k/5)>0:
                tmp = "%s%s"%(self.luomaHash[int(w)],self.luomaHash[w*10])
            else:
#                 print(k,w)
                tmp = "%s%s"%(self.luomaHash[w],self.luomaHash[w*5])
        else:
            
            if int(k/5)>0:
                tmp = "%s%s"%(int(k/5)*self.luomaHash[w*5],int(k%5)*self.luomaHash[w])
            else:
                tmp = "%s%s"%(int(k/5)*self.luomaHash[int(w/2)],int(k%5)*self.luomaHash[w])
        return tmp
