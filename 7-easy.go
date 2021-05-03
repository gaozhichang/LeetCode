/**
【简单|中】
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。

100%,78%
思路：对当前数不断除以10，并对10取余，得到数字。（开始想用字符串方式，特别费劲）
卡点：取数字的顺序就是原来数字的倒序，所以直接在同一个循环操作就行。不用再单独循环。
*/

func reverse(x int) int {
    isfu:=0
    if x<0{
        x = x*-1
        isfu = 1
    }

    // z:=[]int{}
    newx:=0
    for x>0 {
        // z := append(z, x%10)
        zimu:=x%10
        x = x/10
        newx = newx*10 + zimu
    }
    // fmt.Println(z)
    // wei:=1
   
    // for i:=len(z)-1;i>=0;i--{
    //     // fmt.Println(z[i])
    //     newx = wei*z[i] + newx
    //     wei = wei*10
        
    // }
    if isfu==1{
        newx = newx*-1
    }
    if newx < -2147483648 || newx > 2147483647{
        return 0
    }
        

    return newx
}
