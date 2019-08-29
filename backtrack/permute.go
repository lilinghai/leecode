package backtrack

import (
	"fmt"
)

//输出一个不重复数组的所有元素的全排列
//[1,2]	->[1,2],
//1,2,3,4 上一层使用过的元素，下层不能再次使用，需要剪枝
func Permute(nums []int) {
	result:=make([]int,len(nums))
	//backtrack2(nums,0,result)
	used:=make([]bool,len(nums))
	backtrack3(nums,0,used,result)
}

func backtrack2(nums []int,deep int,result []int){
	if len(nums)==0{
		fmt.Println(result)
		return
	}
	for i:=0;i<len(nums);i++{
		result[deep]=nums[i]
		var newNums []int
		newNums=append(newNums,nums[0:i]...)
		newNums=append(newNums,nums[i+1:len(nums)]...)
		backtrack2(newNums,deep+1,result)
	}
}

//使用used数组来标记是否使用过

func backtrack3(nums []int,deep int,used[]bool,result []int){
	if deep>=len(nums){
		fmt.Println(result)
		return
	}
	for i:=0;i<len(nums);i++{
		if !used[i]{
			used[i]=true
			result[deep]=nums[i]
			backtrack3(nums,deep+1,used,result)
			used[i]=false
		}
	}
}
