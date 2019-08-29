package tree

import (
	"container/list"
	"fmt"
)

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

func RecursiveInOrderReverse(root *TreeNode) {
	if root == nil {
		return
	}
	RecursiveInOrderReverse(root.Left)
	fmt.Println(root.Val)
	RecursiveInOrderReverse(root.Right)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//中序遍历判断,如果是bst，中序遍历是有序的，
func IsValidBST(root *TreeNode) bool {
	type TmpNode struct{
		Node *TreeNode
		Visited bool
	}
	lst:=list.New()
	lst.PushBack(TmpNode{root,false})
	var pre *TreeNode=nil
	for lst.Len()!=0&&root!=nil{
		elem:=lst.Back()
		lst.Remove(elem)
		node:=elem.Value.(TmpNode).Node
		visited:=elem.Value.(TmpNode).Visited
		if visited{
			if pre!=nil &&pre.Val>=node.Val{
				return false
			}
			pre=node
			//fmt.Println(node.Val)
		}else{
			if node.Right!=nil{
				lst.PushBack(TmpNode{node.Right,false})
			}
			lst.PushBack(TmpNode{node,true})
			if node.Left!=nil{
				lst.PushBack(TmpNode{node.Left,false})
			}
		}
	}
	return true
}

/**
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先
1.构造map tree
2.链表交点
 */
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode{
	type TmpNode struct{
		Node *TreeNode
		Visited bool
	}
	lst:=list.New()
	lst.PushBack(TmpNode{root,false})
	//每个孩子节点只有一个父节点，树可以看作多个孩子节点指向父节点的链表。这多个链表可以用map表示
	//如 [1,2,3]中序遍历树，可以表示为1->2,3->2,2->nil
	treeMap:=make(map[*TreeNode]*TreeNode)
	treeMap[root]=nil
	for lst.Len()!=0&&root!=nil{
		elem:=lst.Back()
		lst.Remove(elem)
		node:=elem.Value.(TmpNode).Node
		visited:=elem.Value.(TmpNode).Visited
		if visited{
			if node.Right!=nil{
				treeMap[node.Right]=node
			}
			if node.Left!=nil{
				treeMap[node.Left]=node
			}
		}else{
			if node.Right!=nil{
				lst.PushBack(TmpNode{node.Right,false})
			}
			lst.PushBack(TmpNode{node,true})
			if node.Left!=nil{
				lst.PushBack(TmpNode{node.Left,false})
			}
		}
	}
	//构造出孩子节点指向父节点的数据结构之后，找到两个孩子节点的最近公共祖先，可以看作是两个链表相遇的问题
	l1:=p
	l2:=q
	for l1!=l2{
		if l1!=nil{
			l1=treeMap[l1]
		}else{
			l1=q
		}
		if l2!=nil{
			l2=treeMap[l2]
		}else{
			l2=p
		}
	}
	return l1
}

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
//寻找二叉搜索树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	for root!=nil{
		if p.Val>root.Val &&q.Val>root.Val{
			root=root.Right
		}else if p.Val<root.Val &&q.Val<root.Val{
			root=root.Left
		}else{
			return root
		}
	}
	return nil
}