package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
给你一个字符串 time ，格式为 hh:mm（小时：分钟），其中某几位数字被隐藏（用 ? 表示）。

有效的时间为 00:00 到 23:59 之间的所有时间，包括 00:00 和 23:59 。

替换 time 中隐藏的数字，返回你可以得到的最晚有效时间。

 

示例 1：

输入：time = "2?:?0"
输出："23:50"
解释：以数字 '2' 开头的最晚一小时是 23 ，以 '0' 结尾的最晚一分钟是 50 。
示例 2：

输入：time = "0?:3?"
输出："09:39"
示例 3：

输入：time = "1?:22"
输出："19:22"
 

提示：

time 的格式为 hh:mm
题目数据保证你可以由输入的字符串生成有效的时间

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/latest-time-by-replacing-hidden-digits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
这里可以不用递归，直接判断每一位是否位问号，看考官方
 */
func maximumTime(time string) string {
	pos := strings.Index(time,"?")
	switch pos {
	case 4:
		time = strings.Replace(time,"?","9",1)
		break
	case 3:
		time = strings.Replace(time,"?","5",1)
		break
	case 1:
		if string(time[0])=="0" || string(time[0])=="1"{
			time = strings.Replace(time,"?","9",1)
		}else if string(time[0])=="2"{
			time = strings.Replace(time,"?","3",1)
		}else{
			fmt.Println("h1 not match")
		}
		break
	case 0:
		h1,_:=strconv.Atoi(string(time[1]))
		if h1>3{
			time = strings.Replace(time,"?","1",1)
		}else{
			time = strings.Replace(time,"?","2",1)
		}
		break
	default:
		fmt.Println("time not case")
	}
	if strings.Index(time,"?")!=-1{
		time = maximumTime(time)
	}

	return time
}

/**
官方
 */
func maximumTime_pro(time string) string {
	t := []byte(time)
	if t[0] == '?' {
		if '4' <= t[1] && t[1] <= '9' {
			t[0] = '1'
		} else {
			t[0] = '2'
		}
	}
	if t[1] == '?' {
		if t[0] == '2' {
			t[1] = '3'
		} else {
			t[1] = '9'
		}
	}
	if t[3] == '?' {
		t[3] = '5'
	}
	if t[4] == '?' {
		t[4] = '9'
	}
	return string(t)
}


func main()  {
	fmt.Println(maximumTime("00:??"))
	fmt.Println(maximumTime("2?:?0"))
	fmt.Println(maximumTime("0?:3?"))
	fmt.Println(maximumTime("1?:22"))
}
