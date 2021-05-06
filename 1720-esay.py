'''
[简单|难]
未知 整数数组 arr 由 n 个非负整数组成。
经编码后变为长度为 n - 1 的另一个整数数组 encoded ，其中 encoded[i] = arr[i] XOR arr[i + 1] 。例如，arr = [1,0,2,1] 经编码后得到 encoded = [1,2,3] 。
给你编码后的数组 encoded 和原数组 arr 的第一个元素 first（arr[0]）。
请解码返回原数组 arr 。可以证明答案存在并且是唯一的。

【解题】
这是道模拟（重拳出击）题。
根据题目给定的规则，利用如下异或性质从头做一遍即可：
相同数值异或，结果为 0
任意数值与 0 进行异或，结果为数值本身
异或本身满足交换律。

已知 encoded[i-1] = arr[i-1] XOR arr[i]，将等式两边同时「异或」上 arr[i-1]。可得：

encoded[i-1] XOR arr[i-1] = arr[i-1] XOR arr[i] XOR arr[i-1]
结合「性质三」和「性质一」，可化简「右式」得 encoded[i-1] XOR arr[i-1] = arr[i] XOR 0
结合「性质二」，可化简「右式」得 encoded[i-1] XOR arr[i-1] = arr[i]

【卡点】
异或的计算规则不清楚。
技巧：arr[-1]取最后一项

'''


class Solution(object):
    def decode(self, encoded, first):
        """
        :type encoded: List[int]
        :type first: int
        :rtype: List[int]
        """

        arr = [first]
        for one in encoded:
            arr.append(arr[-1]^one)

        return arr


            
