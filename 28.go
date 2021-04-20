/*
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-strstr
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

边缘case:
"",""
"a",""
提交：
100%，28%
*/
func strStr(haystack string, needle string) int {
    if haystack==needle{
        return 0
    }

    n:=len(haystack)
    step:=len(needle)

    if len(haystack)<len(needle){
        return -1
    }
    for i:=0;i<=n-step;i++{
        // fmt.Println(haystack[i:i+step])
        if haystack[i:i+step]==needle{
            return i
        }
            
    }
    return -1
}
