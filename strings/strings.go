package strings

import (
	"fmt"
	"math"
	"strings"
)

// +123abc
//abc123
//-123
//123abc
//00abc
//2147483647 -2147483648
//+-2
//+ 2
//2 3
//+0 123
//0 123
//+ 123
//- 123
/*
核心逻辑 循环遍历
碰到空格，如果空格出现在符号或者数字之后则停止，出现在前面就略过
碰到符号，如果符号出现在符号或者数字之后则停止
碰到数字，累加，累加前要判断最小和最大值
碰到其他，停止
 */
func MyAtoi(str string) int {
	var res  int
	start:=false
	flag:=1
	flagNum:=0
	for i:=0;i<len(str);i++{
		if str[i]==' '{
			if start || flagNum>0{
				break
			}
			continue
		}else if str[i]=='+'{
			if flagNum>0 || start{
				break
			}
			flagNum++
			continue
		}else if str[i]=='-'{
			if flagNum>0 || start{
				break
			}
			flagNum++
			flag=-1
		}else if str[i]<='9' && str[i]>='0'{
			start=true
			value:=(int)(str[i]-'0')*flag
			fmt.Println(res,value)
			if res>math.MaxInt32/10 || (res==math.MaxInt32/10 && value>=math.MaxInt32%10){
				res=math.MaxInt32
				break
			}
			if res<math.MinInt32/10 ||(res==math.MinInt32/10 && value<=math.MinInt32%10){
				res=math.MinInt32
				break
			}
			res=res*10+value
		}else{
			break
		}
	}
	fmt.Println(res)
	return res
}

//["flower","flow","flight"]   fl
//数组字符串的最长公共子串
func LongestCommonPrefix(strs []string) string{
	if len(strs)==0{
		return ""
	}
	prefix:=strs[0]
	for i:=1;i<len(strs);i++{
		for strings.Index(strs[i],prefix)!=0{
			prefix=prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func LongestCommonPrefix2(strs []string) string {
	if len(strs)==0{
		return ""
	}
	for i:=0;i<len(strs[0]);i++{
		for j:=1;j<len(strs);j++{
			if len(strs[j])==i || strs[0][i]!=strs[j][i]{
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

//bm算法，寻找字串在源字符串的位置
/*
	坏字符规则
	acbecb
	bce
	移动2

	acbecb
	  bce
	移动1
 */
func Bm(a string, b string) int {
	bc:=make(map[byte]int)
	for i:=0;i<len(b);i++{
		bc[b[i]]=i
	}

	// i 表示主串与模式串对齐的第一个字符
	for i:=0;i <= len(a) - len(b); {
		j:=0
		for j = len(b) - 1; j >= 0;j--{ // 模式串从后往前匹配
			if a[i+j] != b[j] {
				break
			} // 坏字符对应模式串中的下标是 j
		}
		if j < 0 {
			return i; // 匹配成功，返回主串与模式串第一个匹配的字符的位置
		}
		// 这里等同于将模式串往后滑动 j-bc[(int)a[i+j]] 位
		i = i + (j - bc[a[i+j]]);
	}
	return -1;
}
