package hash

func isAnagram(s string, t string) bool {
	count:=make(map[rune]int)
	for _,value:=range s{
		if _,ok:=count[value];!ok{
			count[value]=0
		}
		count[value]++
	}
	for _,value:=range t{
		if elem,ok:=count[value];!ok||elem==0{
			return false
		}
		count[value]--
	}
	for _,v:=range count{
		if v!=0{
			return false
		}
	}
	return true
}