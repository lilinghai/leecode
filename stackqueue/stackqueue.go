package stackqueue

import "container/list"

//1->2->3->4  使用两个栈来实现一个队列
type MyQueue struct {
	stackIn *list.List
	stackOut *list.List
}


/** Initialize your data structure here. */

func Constructor2() MyQueue {
	return MyQueue{list.New(),list.New()}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.stackIn.PushBack(x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.stackOut.Len()!=0{
		elem:=this.stackOut.Back()
		this.stackOut.Remove(elem)
		return elem.Value.(int)
	}
	for this.stackIn.Len()!=0{
		elem:=this.stackIn.Back()
		this.stackIn.Remove(elem)
		this.stackOut.PushBack(elem.Value.(int))
	}
	if this.stackOut.Len()!=0{
		elem:=this.stackOut.Back()
		this.stackOut.Remove(elem)
		return elem.Value.(int)
	}
	return -1
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.stackOut.Len()!=0{
		elem:=this.stackOut.Back()
		return elem.Value.(int)
	}
	for this.stackIn.Len()!=0{
		elem:=this.stackIn.Back()
		this.stackIn.Remove(elem)
		this.stackOut.PushBack(elem.Value.(int))
	}
	if this.stackOut.Len()!=0{
		elem:=this.stackOut.Back()
		return elem.Value.(int)
	}
	return -1
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.stackOut.Len()==0 &&this.stackIn.Len()==0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
