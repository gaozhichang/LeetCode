/**
【难度】easy
[卡点] 竟然没用取模，自己循环算了。
*/
func chalkReplacer(chalk []int, k int) int {
    var onceUsed int
    //先计算一轮要多少个
    for i:=0;i<len(chalk);i++{
        onceUsed+=chalk[i]
    }
    // var useAll int
    // for useAll<=k{
    //     useAll += onceUsed
    // }
    //计算最后一轮要多少个
    k = k%onceUsed
    // findK := onceUsed-(useAll - k)
    // fmt.Println(onceUsed,useAll,findK)
    //计算最后结果
    for i:=0;i<len(chalk);i++{
        if k<chalk[i]{
            return i 
        }else{
            k = k - chalk[i]
        }
    }
    

    return -1
}
