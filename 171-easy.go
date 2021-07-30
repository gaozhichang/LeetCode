package main
/**
[简单|简单】
【卡点】无
【解答】
1. 按照进制转换的方式做。
2. 但其实strings[i]本身就是 []byte=>[]unit8 即数字
 */
import (
	"fmt"
	"math"
)

func titleToNumber(columnTitle string) int {
	res:=0
	hash:=make(map[string]int)
	hash = map[string]int{"A":1,"B":2,"C":3,"D":4,"E":5,"F":6,"G":7,"H":8,"I":9,"J":10,"K":11,"L":12,"M":13,"N":14,"O":15,"P":16,"Q":17,"R":18,"S":19,"T":20,"U":21,"V":22,"W":23,"X":24,"Y":25,"Z":26}
	for i:=len(columnTitle)-1;i>=0;i--{
		n:=hash[string(columnTitle[i])]
		res = res + n*int(math.Pow(float64(26),float64(len(columnTitle)-1-i)))
		fmt.Println("--->",i,res,string(columnTitle[i]))
	}
	return res
}
func titleToNumberPro(columnTitle string) int {
	res:=0
	//hash:=make(map[string]int)
	//hash = map[string]int{"A":1,"B":2,"C":3,"D":4,"E":5,"F":6,"G":7,"H":8,"I":9,"J":10,"K":11,"L":12,"M":13,"N":14,"O":15,"P":16,"Q":17,"R":18,"S":19,"T":20,"U":21,"V":22,"W":23,"X":24,"Y":25,"Z":26}
	for i:=len(columnTitle)-1;i>=0;i--{
		//columnTitle[i] 这个经典啊，可以直接得到字母A的数字值，在减去A就是偏移了多少
		n:=int(columnTitle[i]-'A'+1)
		res = res + n*int(math.Pow(float64(26),float64(len(columnTitle)-1-i)))
		fmt.Println("--->",i,res,string(columnTitle[i]))
	}
	return res
}

func main()  {
	s:="AB"
	s1 := []byte(s)
	fmt.Println(titleToNumberPro(s))
	fmt.Println(fmt.Sprintf("%T",s))
	fmt.Println(fmt.Sprintf("%T",s1))
	fmt.Println(s[1])
	fmt.Println(s1[1])
}
