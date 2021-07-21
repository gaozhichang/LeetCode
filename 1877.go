package main
/**
[中等|简单]
一个数对 (a,b) 的 数对和 等于 a + b 。最大数对和 是一个数对数组中最大的 数对和 。

比方说，如果我们有数对 (1,5) ，(2,3) 和 (4,4)，最大数对和 为 max(1+5, 2+3, 4+4) = max(6, 5, 8) = 8 。
给你一个长度为 偶数 n 的数组 nums ，请你将 nums 中的元素分成 n / 2 个数对，使得：

nums 中每个元素 恰好 在 一个 数对中，且
最大数对和 的值 最小 。
请你在最优数对划分的方案下，返回最小的 最大数对和 。

【解题】读懂题很难，然后就是排序，让最大和最小组合
【卡点】
1. 读懂题

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimize-maximum-pair-sum-in-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
**/
import (
	"fmt"
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	var res int
	for i:=0;i<n/2;i++{
		new:=nums[i]+nums[n-i-1]
		if res<new{
			res = new
		}
		fmt.Println(new,res)
	}

	return res
}
func main()  {
	mums:=[]int{3,5,2,3}
	res:= minPairSum(mums)
	fmt.Println(res)
}
