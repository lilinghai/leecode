package strings

import "strconv"

//aabbcd
func ZipStr(str string)string{
	var res []byte
	count:=1
	for i:=0;i<len(str);i++{
		if count==1{
			res=append(res,str[i])
		}
		if i+1<len(str) && str[i]==str[i+1]{
			count++
		}
		if i+1<len(str) && str[i]!=str[i+1] && count!=1{
			res=append(res,strconv.Itoa(count)...)
			count=1
		}
		if i+1==len(str) && count!=1{
			res=append(res,strconv.Itoa(count)...)
		}
	}
	return string(res)
}
