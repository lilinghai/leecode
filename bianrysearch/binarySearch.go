package bianrysearch

//[1,3,5,7,8,9,,13]
//数组中有多个key存在时，返回一个随机的key出现的位置
func Search(arr []int,key int)int{
	start:=0
	end:=len(arr)-1
	for end>=start{
		mid:=(start+end)/2
		if arr[mid]==key{
			return mid
		}
		if arr[mid]>key{
			end=mid-1
		}
		if arr[mid]<key{
			start=mid+1
		}
	}
	//fmt.Println(start,end)
	return -1
}

//1,2,2,2,2,3,4,5,6
//找到第一个出现key的索引，或者是最后一个出现key的索引
//此种方法需要考虑数组为空，数组长度时1，和数组长度时2的特殊情况
func Search2(arr []int,key int)int{
	start:=0
	end:=len(arr)-1
	for start+1<end{
		mid:=(start+end)/2
		//根据条件判断如果最后一个则不断的调整start；如果第一个则不断调整end
		if arr[mid]<=key{
			start=mid
		}else{
			end=mid
		}
	}
	//数组长度为2，没有进入循环，如果要寻找最后一个出现key的位置，则需要先判断end
	if len(arr)!=0&&arr[end]==key{
		return end
	}
	if len(arr)!=0&&arr[start]==key{
		return start
	}
	return -1
}

/*
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
输出: true
 */

func SearchMatrix(matrix [][]int,target int)bool{
	if len(matrix)==0 || len(matrix[0])==0{
		return false
	}
	start:=0
	end:=len(matrix)-1
	size:=len(matrix[0])
	line:=-1
	for start<=end{
		mid:=(start+end)/2
		if matrix[mid][0]==target{
			return true
		}
		if matrix[mid][0]>target{
			end=mid-1
		}
		if matrix[mid][0]<target{
			if matrix[mid][size-1]==target{
				return true
			}
			if matrix[mid][size-1]>target{
				line=mid
				break
				//找到了所在的行
			}
			if matrix[mid][size-1]<target{
				start=mid+1
			}
		}
	}
	if line==-1{
		return false
	}
	start=0
	end=size-1
	for start<=end{
		mid:=(start+end)/2
		if matrix[line][mid]==target{
			return true
		}
		if matrix[line][mid]>target{
			end=mid-1
		}
		if matrix[line][mid]<target{
			start=mid+1
		}
	}
	return false
}

/*
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
从左下角开始遍历，如果大于则向右移动，小于则向左移动
 */

func SearchMatrix2(matrix[][]int,target int)bool{
	if len(matrix)==0 || len(matrix[0])==0{
		return false
	}
	row:=len(matrix)-1
	column :=0
	for row>=0 && column <len(matrix[0]){
		if matrix[row][column]==target{
			return true
		}
		if matrix[row][column]>target{
			row--
			continue
		}
		if matrix[row][column]<target{
			column++
		}
	}
	return false
}

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组  [0，1，2，4，5，6，7]  可能变为 [4，5，6，7，8,9,0，1，2,3] )。
7,0,1,2,3,4,5,6
请找出其中最小的元素。
旋转数组的特性1.包含两个有序数组；2.如果旋转，则前面数组最小元素大于后面数组最大元素。
如果旋转，最后一个元素是第二个有序数组的最大值，和这个最大值进行比较来进行二分查找，如果mid大于最后一个元素，说明mid在前面区间，
如果小于最后一个元素，说明mid在后面区间
 */

func SearchReverseArr(arr []int)int{
	start:=0
	end:=len(arr)-1
	for start+1<end{
		mid:=(start+end)/2
		//说明经过了旋转，则前面有序数组最小元素大于后面有序数组最大元素
		if arr[mid]>arr[end]{
			start=mid
		}else{
			end=mid
		}
	}
	if arr[end]<arr[start]{
		return arr[end]
	}
	return arr[start]
}

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组[0,1,2,4,5,6,7] 可能变为[4,5,6,7,8，9，0,1,2，3])。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。
 */

func SearchReverseArr2(arr []int,target int)int{
	start:=0
	end:=len(arr)-1
	for start+1<end{
		mid:=(start+end)/2
		if target>arr[end] {
			if arr[mid] > target {
				end = mid
			}
			if arr[mid] == target {
				return mid
			}
			//arr[mid]小于target分两种情况，一种是在后面区间，一种是在前面区间
			if arr[mid] < target {
				if arr[mid] < arr[end] {
					end = mid
				}else{
					start=mid
				}
			}
		}else{
			if arr[mid]<target{
				start=mid
			}
			if arr[mid]==target{
				return mid
			}
			if arr[mid]>target{
				if arr[mid]>arr[end]{
					start=mid
				}else{
					end=mid
				}
			}

		}
	}
	if arr[start]==target{
		return start
	}
	if arr[end]==target{
		return end
	}
	return -1
}

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
3,3,3,3,3,3,4,5,3,3,3,3   target 4
4在前面区间，arr[mid]=3 和最后一个元素比较，不能确定是在前面区间或是后面区间，所以该题和上题的主要区别就是
在arr[mid]=arr[end]的情况下确定mid在哪个区间
编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
 */
func SearchReverseArr3(arr []int,target int)bool{
	start:=0
	end:=len(arr)-1
	for start+1<end{
		mid:=(start+end)/2
		if target>arr[end] {
			if arr[mid] > target {
				end = mid
			}
			if arr[mid] == target {
				return true
			}
			//arr[mid]小于target分两种情况，一种是在后面区间，一种是在前面区间
			if arr[mid] < target {
				if arr[mid] < arr[end] {
					end = mid
				}else if arr[mid]>arr[end]{
					start=mid
				}else{
					tmp:=mid
					for tmp<end && arr[tmp]==arr[end]{
						tmp++
					}
					if arr[tmp]==arr[end]{
						end=mid
					}else{
						start=mid
					}
				}
			}
		}else{
			if arr[mid]<target{
				start=mid
			}
			if arr[mid]==target{
				return true
			}
			if arr[mid]>target{
				if arr[mid]>arr[end]{
					start=mid
				}else if arr[mid]<arr[end]{
					end=mid
				}else{
					tmp:=mid
					for tmp<end && arr[tmp]==arr[end]{
						tmp++
					}
					if arr[tmp]==arr[end]{
						end=mid
					}else{
						start=mid
					}
				}
			}

		}
	}
	if len(arr)>0&&arr[start]==target{
		return true
	}
	if len(arr)>0&&arr[end]==target{
		return true
	}
	return false
}

/*
给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。
 */
func SquarePerfect(num int)bool{
	start:=0
	end:=num
	for start+1<end{
		mid:=start+(end-start)/2
		if mid*mid>=num{
			end=mid
		}else if mid*mid<num{
			start=mid
		}
	}
	if num==1 || start*start==num || end*end==num{
		return true
	}
	return false
}

/*
实现 int sqrt(int x) 函数。
计算并返回 x 的平方根，其中 x 是非负整数。
由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
sqrt(12)=3
 */

func Sqrt(x int)int{
	start:=0
	end:=x
	for start+1<end{
		mid:=start+(end-start)/2
		if mid*mid>=x{
			end=mid
		}else if mid*mid<x{
			start=mid
		}
	}
	if x==1{
		return 1
	}
	if end*end==x{
		return end
	}
	return start
}

/*
x 的平方根
如果输入参数是任意的浮点数
 */










