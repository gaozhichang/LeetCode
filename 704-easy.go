/*
@Time : 2021/9/6 上午10:22
@Author : gaozhichang
@File : binary-search.go
@des: GoLand
*/
package main

import "fmt"

func search(nums []int, target int) int {
	n:=len(nums)
	if n==0{
		return -1
	}

	hashMap:=make(map[int]int)
	for i,v:=range nums{
		hashMap[v] = i
	}

	var findMid func([]int) int
	findMid = func(nums []int) int{
		fmt.Println(nums)
		if len(nums)==0{
			return -1
		}
		if len(nums)==1 && nums[0]!=target{
			return -1
		}
		mid:=len(nums)/2
		if nums[mid]==target{
			return hashMap[nums[mid]]
		}else if nums[mid]>target{
			return findMid(nums[:mid])
		}else{
			return findMid(nums[mid:])
		}
	}
	return findMid(nums)
}

/**
官方
 */
func search2(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}


func main()  {
	nums:=[]int{-1,0,3,5,9,12}
	res:=search(nums,2)
	fmt.Println(res)
}
