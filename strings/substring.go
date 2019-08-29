package strings

//最大不重复子串的长度
//abcdbdaef
//双指针滑动，借助hash判断是否重复
func LengthOfLongestSubstring(s string) int {
	posMap:=make(map[byte]int)
	var res,j int
	for i:=0;i<len(s);i++{
		if pos,ok:=posMap[s[i]];ok{
			if j<pos+1{
				j=pos+1
			}
		}

		if res<i+1-j{
			res=i+1-j
		}
		posMap[s[i]]=i
	}
	return res
}
