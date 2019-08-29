package stackqueue

import (
	"container/heap"
	"fmt"
)

/*
类需要一个同时接收整数 k 和整数数组nums 的构造器，它包含数据流中的初始元素。每次调用 KthLargest.add，返回当前数据流中第K大的元素
 */
type KthLargest struct {
	minHeap *IntHeap
	k int
}


func Constructor(k int, nums []int) KthLargest {
	minHeap:=&IntHeap{}
	heap.Init(minHeap)
	for i,num:=range nums{
		if i<k{
			heap.Push(minHeap,num)
		}else if num>(*minHeap)[0]{
			heap.Pop(minHeap)
			heap.Push(minHeap,num)
		}
	}
	return KthLargest{minHeap:minHeap,k:k}
}


func (this *KthLargest) Add(val int) int {
	if this.minHeap.Len()<this.k{
		heap.Push(this.minHeap,val)
	}else if val>(*(this.minHeap))[0]{
		heap.Pop(this.minHeap)
		heap.Push(this.minHeap,val)
	}
	fmt.Println(this.minHeap,this.minHeap.Len(),this.k,(*(this.minHeap))[0])
	return (*(this.minHeap))[0]
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
