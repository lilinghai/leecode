package mysort

import "fmt"

//冒泡排序
//6，5，3,4
//5,3,4,6	3,4,5,6		3,4,5,6
func BubbleSort(arr []int){
	for i:=0;i<len(arr);i++{
		flag:=false
		for j:=0;j<len(arr)-i-1;j++{
			if arr[j]>arr[j+1]{
				tmp:=arr[j]
				arr[j]=arr[j+1]
				arr[j+1]=tmp
				flag=true
			}
		}
		if !flag{
			break
		}
	}
}

//插入排序
//6，5，3，4
//5，6，3，4
func InsertSort(arr []int){
	for i:=1;i<len(arr);i++{
		value:=arr[i]
		j:=i
		for ;j>0;j--{
			if value<arr[j-1]{
				arr[j]=arr[j-1]
			}else{
				break
			}
		}
		arr[j]=value
	}
}

//选择排序


//归并排序
func MergeSort(array []int, l int,r int) {

	//如果只有一个元素，那就不用排序了
	if (l >= r) {
		return
	}
	//取中间的数，进行拆分
	m:= (l + r) / 2

	//左边的数不断进行拆分
	MergeSort(array, l, m)

	//右边的数不断进行拆分
	MergeSort(array, m + 1, r)

	//合并
	merge(array, l, m, r);

}

func merge(array []int, l int,m int,r int) {
	leftArr:=make([]int,m-l+1)
	rightArr:=make([]int,r-m)
	var li,ri int
	for k:=l;k<r+1;k++{
		if k<=m{
			leftArr[li]=array[k]
			li++
		}else{
				rightArr[ri]=array[k]
				ri++
		}
	}
	var i,j int

	//比较这两个数组的值，哪个小，就往数组上放
	//2,6,10,1,7,8
	for i<len(leftArr)&&j<len(rightArr) {
		if leftArr[i] <= rightArr[j] {
			array[l] = leftArr[i]
			i++
		} else {
			array[l] = rightArr[j]
			j++
		}
		l++
	}

	//如果左边的数组还没比较完，右边的数都已经完了，那么将左边的数抄到大数组中(剩下的都是大数字)
	for i<len(leftArr) {
		array[l]=leftArr[i]
		i++
		l++
	}
	//如果右边的数组还没比较完，左边的数都已经完了，那么将右边的数抄到大数组中(剩下的都是大数字)
	for j<len(rightArr) {
		array[l]=rightArr[j]
		j++
		l++
	}
	fmt.Println(array)
}

//[7 3 10 6 8 1 5]
//选择一个比较点，所有大于比较点的放在右边，小于的放在左边
//5->3,1,5,6,8,7,10
//3，1		6，8，7，10-> 1,3	6,8,7,10
//1,3	6,8,7
func QuickSort(arr[]int,start int ,end int){
	if start>=end{
		return
	}
	//获取分区点
	p:=partition(arr,start,end)
	//fmt.Println(arr,p)
	QuickSort(arr,start,p-1)
	QuickSort(arr,p+1,end)
}

func partition(arr[]int,start int,end int)int{
	base:=arr[end]
	//双指针，j指向第一个大于左边的元素
	j:=start
	for i:=start;i<end;i++{
		if arr[i]<base{
			tmp:=arr[i]
			arr[i]=arr[j]
			arr[j]=tmp
			j++
		}
	}
	arr[end]=arr[j]
	arr[j]=base
	return j
}

//[7 3 10 6 8 1 5]  3	1,3,5,6,7,8,10
func KthLargest(arr[]int,start,end,k int)int{
	if start>=end{
		return arr[start]
	}
	p:=partition(arr,start,end)
	if p+1==k{
		return arr[p]
	}else if p+1>k{
		return KthLargest(arr,start,p-1,k)
	}else{
		return KthLargest(arr,p+1,end,k)
	}
}

//1,9,9,5,5,7,7,7,7,6,1 排序，数字集中在0-9之间
/*
0，1，2，3，4，5，6，7，8，9		数组下标
0，2，0，0，0，2，1，4，0，2		数组元素
如果排序的仅仅是数字，则直接回放计数数组就可以
1，1，5，5
 */
func ScoreSort(arr[]int){

}
