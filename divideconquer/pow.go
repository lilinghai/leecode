package divideconquer

import "fmt"

//基于x^10*x^10=x^20 已经计算出x^10之后，不需要再计算一遍，只需要乘积就可以
//使用分治法，分而治之
func myPow(x float64, n int) float64 {
	//递归退出的条件，n相当于递归的层级，到达最后一层需要退出
	if n==0{
		return 1
	}
	if n<0{
		return 1/myPow(x,-n)
	}
	//每次递归需要改变层级，也就是n，同时也要改变参数
	if n%2==0{
		return myPow(x*x,n/2)
	}else{
		return myPow(x*x,n/2)*x
	}
}

func myPow2(x float64, n int) float64 {
	if n<0{
		x=1/x
		n=-n
	}
	var res float64=1
	for n!=0{
		if n%2==1{
			res*=x
		}
		x*=x
		n/=2
	}
	return res
}

func MyPow(x float64,n int){
	fmt.Println(myPow(x,n))
	fmt.Println(myPow2(x,n))

}

