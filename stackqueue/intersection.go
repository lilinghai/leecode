package stackqueue

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//寻找A，B两个链表的最近的交点，交点之后的链表共享
type ListNode struct {
	  Val int
	  Next *ListNode
}
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	la:=headA
	lb:=headB
	var c1,c2 int
	for la!=lb{
		if la!=nil{
			la=la.Next
		}else{
			la=headB
			c1++
		}
		if lb!=nil{
			lb=lb.Next
		}else{
			lb=headA
			c2++
		}
	}
	fmt.Println(c1,c2)
	return la
}
