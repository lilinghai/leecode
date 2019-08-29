package hash

import "sort"

/*
判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
	给定数组 nums = [-1, 0, 1, 2, -1, -4]
	满足要求的三元组集合为：
	[
	 [-1, 0, 1],
	 [-1, -1, 2]
	]
-4,-1,-1,0,1,2
 */
//-4,-1,-1,0,1,2,2,3,4,5
//-4,-1,5	-4,0,4, -4,1,3	-4,2,2

func ThreeSum(nums []int, target int) [][]int {
	var res[][]int

	for i:=0;i<len(nums);i++{
		numsMap:=make(map[int]int)
		for j:=0;j<i+1;j++{
			if _,ok:=numsMap[-nums[i]-nums[j]];ok{
				res=append(res,[]int{nums[i],nums[j],-nums[i]-nums[j]})
			}else{
				numsMap[-nums[i]-nums[j]]=1
			}
		}
	}
	return res
}

func ThreeSum2(nums []int,target int)[][]int{
	sort.Ints(nums)
	var res[][]int
	for i:=0;i<len(nums);i++{
		if i>0&&nums[i]==nums[i-1]{
			continue
		}
		l:=i+1
		r:=len(nums)-1
		for l<r{
			sum:=nums[i]+nums[l]+nums[r]
			if sum==target{
				res=append(res,[]int{nums[i],nums[l],nums[r]})
				for l<r&&nums[l]==nums[l+1]{
					l++
				}
				for l<r&&nums[r]==nums[r-1]{
					r--
				}
				l++
				r--
			}
			if sum>target{
				r--
			}
			if sum<target{
				l++
			}
		}
	}
	return res
}

func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i:=0;i<len(nums);i++{
		if i>0&&nums[i]==nums[i-1]{
			continue
		}
		for j:=i+1;j<len(nums);j++{
			if j>i+1&&nums[j]==nums[j-1]{
				continue
			}
			l:=j+1
			r:=len(nums)-1
			for l<r{
				sum:=nums[i]+nums[j]+nums[l]+nums[r]
				if sum==0{
					res=append(res,[]int{nums[i],nums[j],nums[l],nums[r]})
					for l<r&&nums[l]==nums[l+1]{
						l++
					}
					for l<r&&nums[r]==nums[r-1]{
						r--
					}
					l++
					r--
				}
				if sum>0{
					r--
				}
				if sum<0{
					l++
				}
			}
		}
	}
	return res
}