package bitops

import "fmt"

/*
一个整型数组里除了两个数字之外，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
*/

/*
能够把袁术数组分为两个子数组，每个子数组中包含一个只出现一次的数组，其他数字出现两次
这两个数字不一样，所以亦或结果部位0，结果至少有一位是1，找到第一个为1的位置，记为第N位。
以第N位是否为1作为便准把原数组分为两个子数组。
11001 是最后的亦或值，第一个（随便一位即可）出现1的位置，如最后一位，这说明该位 两个值不一样
以最后以为是否为1来分割数组，则两个奇数必然分割到不同的数组中，如果数字相同，则必然不会分隔到不同的数组中
*/

func FindNum(arr []int){
	res:=0
	for i:=0;i<len(arr);i++{
		res^=arr[i]
	}
	res2:=res
	res3:=res
	var pos uint=0
	for res2&1==0{
		pos++
		res2>>=1
	}
	for i:=0;i<len(arr);i++{
		if (arr[i]>>pos) &1 ==1{
			res3^=arr[i]
		}
	}
	fmt.Println(res3,res3^res)
}


/*
给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在众数。
输入: [2,2,1,1,1,2,2]
输出: 2
 */
/*
1.构建map，每次遍历的时候比较次数是否大于n/2，如果大于就是众数
2. 摩尔投票，每次从序列选择两个不同的数字（抵消）删除，剩下最后一个或几个相同的数字就是出现次数大于总数一半的
 */
func MajorityElement(nums []int) int {
	res:=nums[0]
	count:=1
	for i:=1;i<len(nums);i++{
		fmt.Println(res,nums[i],count)
		if count==0{
			res=nums[i]
			count=1
		}else{
			if res==nums[i]{
				count++
			}else{
				count--
			}
		}
	}
	return res
}

/*
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
说明：解集不能包含重复的子集。
输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

 */
/*
1.回溯，如何剪枝
2.2^len(nums)个结果，000，001，……，111，比较每一位，如果为1则选中数组对应的位置的元素
 */
func Subsets(nums []int) [][]int {
	var res [][]int
	count:=1<<uint(len(nums))
	for i:=0;i<count;i++{
		var arr[]int
		for j:=0;j<len(nums);j++{
			if i&(1<<uint(j))!=0{
				arr=append(arr,nums[j])
			}
		}
		res=append(res,arr)
	}
	return res
}

/*
所有 DNA 由一系列缩写为 A，C，G 和 T 的核苷酸组成，例如：“ACGAATTCCG”。在研究 DNA 时，识别 DNA 中的重复序列有时会对研究非常有帮助。
编写一个函数来查找 DNA 分子中所有出现超过一次的10个字母长的序列（子串）。
输入: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
输出: ["AAAAACCCCC", "CCCCCAAAAA"]
 */
/*
1.map记录连续长度为10的子串对应出现的次数
2.
 */

















