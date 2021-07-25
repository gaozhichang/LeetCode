func restoreArray(adjacentPairs [][]int) []int {
    var res []int
    hash:=make(map[int][]int)
    //利用map找相邻元素：map[1:[2] 2:[1 3] 3:[4 2] 4:[3]]
    for _,v:=range adjacentPairs{
        l,r:=v[0],v[1]
        hash[l] = append(hash[l],r)
        hash[r] = append(hash[r],l)
    }
    // fmt.Println(hash)
    //找到第一个元素,和第二个元素
    for k,item:=range hash{
        if len(item)==1{//如果周边为1，所以是首元素或尾元素
            res = append(res,k)
            res = append(res,item[0])
            delete(hash,k)
            break
        }
    }
    // fmt.Println(hash)
    // fmt.Println(res)
    for i:=1;len(hash)>1;i++{
        // fmt.Println(res,i,hash)
        find:=hash[res[i]]
        if _,ok:=hash[find[0]];ok{
            res = append(res,find[0])
        }else{
            res = append(res,find[1])
        }
        delete(hash,res[i])
        
    }
    


    
    
    return res
}
