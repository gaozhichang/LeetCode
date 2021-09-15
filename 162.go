/**
【中】难-
【卡点】
1. 用O(n)的方法，没有判断好边界。没有想到最简单的找最大值就是峰值的方法。
2. 局部峰值，这种方法没有想到理论：俗话说「人往高处走，水往低处流」。如果我们从一个位置开始，不断地向高处走，那么最终一定可以到达一个峰值位置。
   然后就是，边界问的判断。案例解法给了特别好的，方法，封装get负责找数组的值，如果没有返回 math.Minint64 即最小值。
*/

//O(n) 这不就是找最大值吗。
func findPeakElement2(nums []int) (idx int) {
    for i,v:=range nums{
        if v>nums[idx]{
            idx=i
        }
    }
    fmt.Println(math.MinInt64)
    return 
}
//O(logn) 人往高处走，水往低处流，沿着一个方向往上走必有最大值。
func findPeakElement(nums []int) int{
    n := len(nums)
    get:=func(i int) int{
        if i==-1 || i==n{
            return math.MinInt64
        }
        return nums[i]
    }

    left:=0
    right:=n-1
    for {
        mid := (right-left)/2 + left
        //如果已经是最大值，直接返回。
        if get(mid-1)<get(mid) && get(mid)>get(mid+1){
            return mid
        }
        //如果左边大，往左边走
        if get(mid-1)>get(mid){
            right = mid-1
        }else{
            //否则就是右边大，因为如果左右都小，不就第一个if就满足了吗
            left = left + 1
        }
    }
    
}
