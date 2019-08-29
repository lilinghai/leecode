package backtrack

//1,2
func SubSet(nums[]int,i int,subset* []int,res*[]*[]int){
	if i>=len(nums){
		return
	}
	subset=new([]int)
	*res=append(*res,subset)
	SubSet(nums,i+1,subset,res)
	*subset=append(*subset,nums[i])
	SubSet(nums,i+1,subset,res)
}

//通过位运算找到数组每一位的状态（是否存在）
func SubSet2(nums[]int)[][]int{
	size:=uint(len(nums))
	num:=1<<size
	res:=make([][]int,num)
	for i:=0;i<num;i++{
		var tmp []int
		for j:=uint(0);j<size;j++{
			if i&(1<<j)!=0{
				tmp=append(tmp,nums[j])
			}
		}
		res[i]=tmp
	}
	return res
}


