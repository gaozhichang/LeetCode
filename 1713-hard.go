package main

/**
[难|难]
【卡点】能想到是用lcs做。但是lcs又太复杂。
【解题】
1. 当其中一个数组元素各不相同时，最长公共子序列问题（LCS）可以转换为最长上升子序列问题（LIS）进行求解。
2. 巧妙运用Sort.SearchInts(arr,o) 返回要找字符的应该位置。

【题目】
给你一个数组 target ，包含若干 互不相同 的整数，以及另一个整数数组 arr ，arr 可能 包含重复元素。

每一次操作中，你可以在 arr 的任意位置插入任一整数。比方说，如果 arr = [1,4,1,2] ，那么你可以在中间添加 3 得到 [1,4,3,1,2] 。你可以在数组最开始或最后面添加整数。

请你返回 最少 操作次数，使得 target 成为 arr 的一个子序列。

一个数组的 子序列 指的是删除原数组的某些元素（可能一个元素都不删除），同时不改变其余元素的相对顺序得到的数组。比方说，[2,7,4] 是 [4,2,3,7,2,1,4] 的子序列（加粗元素），但 [2,4,2] 不是子序列。

 

示例 1：

输入：target = [5,1,3], arr = [9,4,2,3,4]
输出：2
解释：你可以添加 5 和 1 ，使得 arr 变为 [5,9,4,1,2,3,4] ，target 为 arr 的子序列。
示例 2：

输入：target = [6,4,8,1,3,2], arr = [4,7,6,2,3,8,6,1]
输出：3
 

提示：

1 <= target.length, arr.length <= 105
1 <= target[i], arr[i] <= 109
target 不包含任何重复元素。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-operations-to-make-a-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
import (
	"fmt"
	"sort"
)

func minOperations(target []int, arr []int) int {
	res := 0
	keys:=make(map[int]int)
	for i,v:=range target{
		keys[v] = i
	}
	fmt.Println(keys)
	var find []int
	for _,v:=range arr{
		if idx,has:=keys[v];has{
			//fmt.Println(v,idx)
			//searchInts是找给定元素的位置，如果没有也回返回应该在的位置：
			// fmt.Println(sort.SearchInts([]int{0,5}, 4))=1；
			// fmt.Println(sort.SearchInts([]int{0,5}, 6))=2；
			// fmt.Println(sort.SearchInts([]int{0,5}, 6))=1
			if p:=sort.SearchInts(find,idx);p<len(find){
				find[p] = idx
			}else{
				find = append(find, idx)
			}
		}
	}
	res = len(target)-len(find)
	return res
}

func minOperationsPro(target, arr []int) int {
	n := len(target)
	pos := make(map[int]int, n)
	for i, val := range target {
		pos[val] = i
	}
	fmt.Println(pos)
	d := []int{}
	for _, val := range arr {
		if idx, has := pos[val]; has {
			if p := sort.SearchInts(d, idx); p < len(d) {
				d[p] = idx
				fmt.Println(d,p,idx,val)
			} else {
				d = append(d, idx)
				fmt.Println(d)
			}

		}

	}
	return n - len(d)
}


func main()  {
	target:=[]int{5,1,3}
	arr:=[]int{9,3,2,3,4}
	fmt.Println(minOperations(target,arr))

	target = []int{6,4,8,1,3,2}
	arr = []int{4,7,6,2,3,8,6,1}
	fmt.Println(minOperations(target,arr))

	fmt.Println("///")
	fmt.Println(sort.SearchInts([]int{0,5}, 4))
	fmt.Println(sort.SearchInts([]int{0,5}, 6))

}
