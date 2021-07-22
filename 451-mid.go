/**
[中|易]
*/
func frequencySort(s string) string {
    n:=len(s)
    m:=make(map[string]int)
    for i:=0;i<n;i++{
        m[string(s[i])]++
    }
    fmt.Println(m)
    res :=""
    value := make([]int,len(m))
	i := 0
	for _,val := range m{
		value[i]=val
		i++
	}
	sort.Ints(value)
    fmt.Println(value)
	for i,_ := range value{
		for k,v := range m{
            // fmt.Println(i,k,v)
			if v== value[i]{
                // fmt.Println(k)
				res = fmt.Sprintf("%s%s",strings.Repeat(k,v),res)
                delete(m, k)
			}
		}
        // break
	}

    return res
}
