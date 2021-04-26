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
