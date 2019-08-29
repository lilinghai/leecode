package backtrack

import "fmt"

/*
{"a":2,"b":3,"c":2}
可以组合成aabbbcc，abbacc等情况，把所有的情况列举出来
 */

func MapConbine(data map[string]int){
	length:=0
	for _,v:=range data{
		length+=v
	}
	res:=make([]string,length)
	backtrack6(data,res,length,0)
}

func backtrack6(data map[string]int,res[]string,length,index int){
	if index==length{
		fmt.Println(res)
		return
	}
	for k,v:=range data{
		if v>0{
			data[k]--
			res[index]=k
			backtrack6(data,res,length,index+1)
			data[k]++
		}
	}
}
