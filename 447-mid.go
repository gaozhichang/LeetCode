/**
[mid]难-
[卡点]没想出来怎么计算，以及排列组合公式：从m个点中取出2个点的排列 m*(m-1)
*/

func numberOfBoomerangs(points [][]int) int {
    var ans int
    for i,point:=range points{
        //每个点计算一次，就是中心点。不能全局计算（因为定义的是到当前点距离，是相对的，如果全局就是多个点到多个点距离相同了。)
        cnt := make(map[int]int)
        // fmt.Println(i, point,cnt)
        for j,one:=range points{
            //除了同一个节点，其他的都要计算
            if i!=j{
                dis := math.Pow(float64(one[0]-point[0]),2)+math.Pow(float64(one[1]-point[1]),2)
                cnt[int(dis)]++
            }
        }
        //距离相同的点中取2个排列组合公式：m*(m-1)
        for _, m := range cnt {
            ans += m * (m - 1)
        }
        fmt.Println(cnt)
    }
    

    return ans
}
