package backtrack

import (
	"fmt"
	"strconv"
)

/*
一个字符串被5个竖线分割为6段，两个竖线不能紧贴，每个分段数字小于600，求所有满足条件的结果
如123|41|025|43|2|1  	123，41，25，43，2，1
123|4|5|6|7|8
123|45
 */
func StrSegment(str string){
	result:=make([]string,6)
	backtrack4(str,0,result)
}

func backtrack4(str string,deep int,result []string){
	if deep>=6{
		fmt.Println(result)
		return
	}
	for i:=0;i<len(str);i++{
		value,_:=strconv.Atoi(str[0:i+1])
		//小于600，且剩余的元素可以被剩下的竖线分割
		if value<600 && len(str)-i>5-deep{
			//最后一层需要特殊处理
			if deep==5 &&i+1==len(str){
				result[deep]=str
				backtrack4("",deep+1,result)
			}
			if deep!=5{
				result[deep] = str[0 : i+1]
				backtrack4(str[i+1:len(str)], deep+1, result)
			}
		}
	}
}