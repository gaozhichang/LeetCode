/**
[简单|简单中]
卡点：递归可行，但是超时。
解题：用动态规划，滑动数组解决，负载度 O(n)

泰波那契序列 Tn 定义如下： 

T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2

给你整数 n，请返回第 n 个泰波那契数 Tn 的值。

 

示例 1：

输入：n = 4
输出：4
解释：
T_3 = 0 + 1 + 1 = 2
T_4 = 1 + 1 + 2 = 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-th-tribonacci-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func tribonacci(n int) int {
    if n==0{
        return 0
    }else if n<3{
        return 1
    }
    n3,n2,n1,res := 0,1,1,0
    for i:=3;i<=n;i++{
        res = n3+n2+n1
        n3 = n2
        n2 = n1
        n1 = res
        
        
    }
    return res
}
