package array

//使用分治法，每个部分分别求最大和最小值
//1,7,3,4,9,0  求最大值和最小值，需要比较2n次，优化后比较n次
//f(n)=f(n/2)*2+2  f(2)=1 f(4)=6  f(8)=14	f(16)=30
//f(n)=2*f(n/2)+2=2*2*f(n/4)+2*2+2=2*2*2*f(n/8)+2*2*2+2*2+2=2^m*f(n/2^m)+2^m+2^m-1 ...+ ...2=2^m*f(n/2^m)+2^m+1 -2
//f(n/2^m)=f(2)=1   n=2^m+1 带入上式
//f(n)=2^m*f(2)+2^m+1 -2=2^m + 2^m+1 -2=n/2+n-2=1.5n-2
func MaxMin(arr []int)(int,int){
	if len(arr)==1{
		return arr[0],arr[0]
	}
	if len(arr)==2{
		if arr[0]<arr[1]{
			return arr[1],arr[0]
		}
		return arr[0],arr[1]
	}
	max1,min1:=MaxMin(arr[:len(arr)/2])
	max2,min2:=MaxMin(arr[len(arr)/2:])
	max:=max1
	if max1<max2{
		max=max2
	}
	min:=min1
	if min1>min2{
		min=min2
	}
	return max,min
}
