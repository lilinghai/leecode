package backtrack

import "fmt"

//2-9中所有英文字母的组合

var numLetter map[byte]string

func LetterCombinations(digits string)[]string {
	var res []string
	numLetter=map[byte]string{
		'2':"abc",
		'3':"def",
		'4':"ghi",
		'5':"jkl",
		'6':"mno",
		'7':"pqrs",
		'8':"tuv",
		'9':"wxyz",
	}
	backtrack(digits,0,make([]byte,len(digits)),&res)
	//fmt.Println(res)
	return res
}

func backtrack(digits string,deep int,result []byte,res *[]string){
	if len(digits)==0{
		//回溯了一个完整路径
		fmt.Println(string(result))
		*res=append(*res,string(result))
		return
	}
	for i:=0;i<len(numLetter[digits[0]]);i++{
		result[deep]=numLetter[digits[0]][i]
		backtrack(digits[1:len(digits)],deep+1,result,res)
	}
}
