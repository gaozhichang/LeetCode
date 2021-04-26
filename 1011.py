/*
传送带上的包裹必须在 D 天内从一个港口运送到另一个港口。

传送带上的第 i 个包裹的重量为 weights[i]。每一天，我们都会按给出重量的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。

返回能在 D 天内将传送带上的所有包裹送达的船的最低运载能力。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
import math 
def shipWithinDays(weights, D):
        """
        :type weights: List[int]
        :type D: int
        :rtype: int
        """
        l = len(weights)
        total  = sum([i for i in weights])
        maxweight  = max([i for i in weights])
        print("total",total)
        print("maxweight",maxweight)
        minres = math.ceil(total/D)
        minres = max(minres,maxweight)
        print("minres",minres)
        res = minres
        n = 0
        while 1:
            full = 0
            for i in range(l):
                full = full+weights[i]
                if i<l-1 and full+weights[i+1]>res:
#                     print(res,weights[i])
                    full = 0
                    n=n+1
                
            n = n + 1
            print("res:",res,"n:",n)
            if n>D:
                n = 0
                res = res +1
            else:
                break
        res = int(res)
        return res
