package backtrack

//返回1……n 所有k个数的组合
//n=4,k=2	12,13,14,23,24,34

//如果对于如果剪枝不太清楚，可以把回溯的结果图画出来，有助于找到如何剪枝
func Combine(n int, k int) [][]int {
	var res [][]int
	result:=make([]int,k)
	backtrack5(n,k,1,0,result,&res)
	return res
}

func backtrack5(n int,k int,start int,deep int,result []int,res *[][]int){
	if deep==k{
		//fmt.Println(result)
		tmp:=make([]int,k)
		copy(tmp,result)
		*res=append(*res,tmp)
		return
	}
	for i:=start;i<=n;i++{
		result[deep]=i
		backtrack5(n,k,i+1,deep+1,result,res)
	}
}
