package mysort

//链表实现插入排序,插入到已经排好序的部分，从链表删除一个元素，然后插入到已经排好序的链表
//6,2,9,4,10,1,8		2，6，9，4，10，1，8
type ListNode struct {
  	Val int
  	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	cur:=head
	//需要返回最小的节点
	res:=head
	var pre *ListNode
	for cur!=nil{
		var pre2 *ListNode
		cur2:=res
		for cur2!=cur{
			if cur.Val<cur2.Val{
				//第一个节点是空，说明是最小的节点
				if pre2==nil{
					pre.Next=cur.Next
					cur.Next=cur2
					res=cur
					cur=pre
					break
				}else{
					pre.Next=cur.Next
					pre2.Next=cur
					cur.Next=cur2
					cur=pre
					break
				}
			}
			pre2=cur2
			cur2=cur2.Next
		}
		pre=cur
		cur=cur.Next
	}
	return res
}
